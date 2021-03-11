package square

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	baseurl := i.c.endpoint("catalog/list")
	q := baseurl.Query()
	stringTypes := make([]string, len(i.types))
	for i, oneType := range i.types {
		stringTypes[i] = string(oneType)
	}
	err := encoder.Encode(&struct {
		Types  string `schema:"types,omitempty"`
		Cursor string `schema:"cursor,omitempty"`
	}{
		Types:  strings.Join(stringTypes, ","),
		Cursor: i.cursor,
	}, q)
	if err != nil {
		return i.setError(fmt.Errorf("error populating url paramters: %w", err))
	}
	baseurl.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", baseurl.String(), nil)
	if err != nil {
		return i.setError(fmt.Errorf("error generating new request: %w", err))
	}

	req = req.WithContext(i.ctx)

	resp, err := i.c.httpClient.Do(req)
	if err != nil {
		return i.setError(fmt.Errorf("Error with square http request: %w", err))
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
		return i.setError(fmt.Errorf("Error reading response body: %w", err))
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
		return i.setError(fmt.Errorf("Error unmarshalling json response: %w", err))
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
