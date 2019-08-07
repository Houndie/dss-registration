package square

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type ListCatalogIterator interface {
	Value() *CatalogObject
	Error() error
	Next() bool
}

type listCatalogIterator struct {
	types  []CatalogObjectType
	cursor string

	done bool
	idx  int

	values []*CatalogObject
	err    error

	c   *Client
	ctx context.Context
}

func (i *listCatalogIterator) setError(err error) bool {
	i.values = nil
	i.err = err
	return false
}

func (i *listCatalogIterator) Value() *CatalogObject {
	return i.values[i.idx]
}

func (i *listCatalogIterator) Error() error {
	return i.err
}

func (i *listCatalogIterator) Next() bool {
	i.idx = i.idx + 1
	if i.idx < len(i.values) {
		return true
	}

	if i.done {
		return false
	}

	baseurl, err := url.Parse("https://connect.squareup.com/v2/catalog/list")
	if err != nil {
		return i.setError(errors.Wrap(err, "Error parsing url"))
	}

	q := baseurl.Query()
	stringTypes := make([]string, len(i.types))
	for i, oneType := range i.types {
		stringTypes[i] = string(oneType)
	}
	err = encoder.Encode(&struct {
		Types  string `schema:"types,omitempty"`
		Cursor string `schema:"cursor,omitempty"`
	}{
		Types:  strings.Join(stringTypes, ","),
		Cursor: i.cursor,
	}, q)
	if err != nil {
		return i.setError(errors.Wrap(err, "error populating url paramters"))
	}
	baseurl.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", baseurl.String(), nil)
	if err != nil {
		return i.setError(errors.Wrap(err, "error generating new request"))
	}

	req = req.WithContext(i.ctx)

	resp, err := i.c.httpClient.Do(req)
	if err != nil {
		return i.setError(errors.Wrap(err, "Error with square http request"))
	}
	defer resp.Body.Close()

	var errMsg error
	if resp.StatusCode != http.StatusOK {
		errMsg = unexpectedCodeError(resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if errMsg != nil {
			return i.setError(errMsg)
		}
		return i.setError(errors.Wrap(err, "Error reading response body"))
	}
	respJson := struct {
		Errors  []*Error         `json:"errors,omitempty"`
		Cursor  string           `json:"cursor,omitempty"`
		Objects []*CatalogObject `json:"objects,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		if errMsg != nil {
			return i.setError(errMsg)
		}
		return i.setError(errors.Wrap(err, "Error unmarshalling json response"))
	}
	if len(respJson.Errors) != 0 {
		return i.setError(&ErrorList{respJson.Errors})
	}
	if errMsg != nil {
		return i.setError(errMsg)
	}

	if len(respJson.Objects) == 0 {
		return false
	}
	i.values = respJson.Objects
	i.idx = 0
	if respJson.Cursor == "" {
		i.done = true
	} else {
		i.cursor = respJson.Cursor
	}
	return true
}

func (c *Client) ListCatalog(ctx context.Context, types []CatalogObjectType) ListCatalogIterator {
	return &listCatalogIterator{
		types:  types,
		cursor: "",
		done:   false,
		idx:    -1,
		values: nil,
		err:    nil,
		c:      c,
		ctx:    ctx,
	}
}
