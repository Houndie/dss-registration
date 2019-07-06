package square

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

func (c *Client) ListLocations() ([]*Location, error) {
	resp, err := c.httpClient.Get("https://connect.squareup.com/v2/locations")
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
