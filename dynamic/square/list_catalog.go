package square

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func (c *Client) ListCatalog(types []string) ([]*CatalogObject, error) {
	baseurl, err := url.Parse("https://connect.squareup.com/v2/catalog/list")
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing url")
	}
	if len(types) != 0 {
		q := baseurl.Query()
		q.Set("types", strings.Join(types, ","))
		baseurl.RawQuery = q.Encode()
	}

	cursor := ""
	objects := []*CatalogObject{}
	for {
		url := baseurl
		if cursor != "" {
			q := url.Query()
			q.Set("cursor", cursor)
			url.RawQuery = q.Encode()
		}
		resp, err := c.httpClient.Get(url.String())
		if err != nil {
			return objects, errors.Wrap(err, "Error with square http request")
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return objects, errors.Wrap(err, "Error reading response body")
		}
		respJson := struct {
			Errors  []*Error         `json:"errors,omitempty"`
			Cursor  string           `json:"cursor,omitempty"`
			Objects []*CatalogObject `json:"objects,omitempty"`
		}{}
		err = json.Unmarshal(bytes, &respJson)
		if err != nil {
			return objects, errors.Wrap(err, "Error unmarshalling json response")
		}
		if len(respJson.Errors) != 0 {
			return objects, &ErrorList{respJson.Errors}
		}
		objects = append(objects, respJson.Objects...)
		if respJson.Cursor == "" {
			break
		}
		cursor = respJson.Cursor
	}
	return objects, nil
}
