package square

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) ListLocations(ctx context.Context) ([]*Location, error) {
	req, err := http.NewRequest("GET", c.endpoint("locations").String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating new request")
	}
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Error listing locations")
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading response body")
	}

	respJson := struct {
		Errors    []*Error    `json:"errors,omitempty"`
		Locations []*Location `json:"locations,omitempty"`
	}{}

	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling json response")
	}
	if len(respJson.Errors) != 0 {
		return nil, &ErrorList{respJson.Errors}
	}
	return respJson.Locations, nil
}
