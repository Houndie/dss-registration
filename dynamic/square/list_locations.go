package square

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListLocations(ctx context.Context) ([]*Location, error) {
	req, err := http.NewRequest("GET", c.endpoint("locations").String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %w", err)
	}
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error listing locations: %w", err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	respJson := struct {
		Errors    []*Error    `json:"errors,omitempty"`
		Locations []*Location `json:"locations,omitempty"`
	}{}

	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling json response: %w", err)
	}
	if len(respJson.Errors) != 0 {
		return nil, &ErrorList{respJson.Errors}
	}
	return respJson.Locations, nil
}
