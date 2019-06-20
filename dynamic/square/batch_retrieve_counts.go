package square

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *Client) BatchRetrieveInventoryCounts(catalogObjectIds, locationIds []string, updatedAfter *time.Time) ([]*InventoryCount, error) {
	url, err := url.Parse("https://connect.squareup.com/v2/inventory/batch-retrieve-counts")
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing url")
	}
	body := struct {
		CatalogObjectIds []string `json:"catalog_object_ids"`
		LocationIds      []string `json:"location_ids"`
		UpdatedAfter     string   `json:"updated_after"`
		Cursor           string   `json:"cursor"`
	}{
		CatalogObjectIds: catalogObjectIds,
		LocationIds:      locationIds,
	}
	if updatedAfter != nil {
		body.UpdatedAfter = updatedAfter.String()
	}

	counts := []*InventoryCount{}
	for {
		jsonBody, err := json.Marshal(&body)
		if err != nil {
			return counts, errors.Wrap(err, "error marshaling request body")
		}
		buf := bytes.NewBuffer(jsonBody)

		resp, err := c.httpClient.Post(url.String(), "application/json", buf)
		if err != nil {
			return counts, errors.Wrap(err, "Error with square http request")
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return counts, errors.Wrap(err, "Error reading response body")
		}
		respJson := struct {
			Errors []*Error          `json:"errors"`
			Cursor string            `json:"cursor"`
			Counts []*InventoryCount `json:"counts"`
		}{}
		err = json.Unmarshal(bytes, &respJson)
		if err != nil {
			return counts, errors.Wrap(err, "Error unmarshalling json response")
		}
		if len(respJson.Errors) != 0 {
			return counts, &ErrorList{respJson.Errors}
		}
		counts = append(counts, respJson.Counts...)
		if respJson.Cursor == "" {
			break
		}
		body.Cursor = respJson.Cursor
	}
	return counts, nil
}
