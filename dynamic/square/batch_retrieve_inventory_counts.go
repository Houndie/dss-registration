package square

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BatchRetrieveInventoryCountsIterator interface {
	Value() *InventoryCount
	Error() error
	Next() bool
}

type batchRetrieveInventoryCountsIterator struct {
	catalogObjectIDs []string
	locationIDs      []string
	updatedAfter     *time.Time
	cursor           string

	counts []*InventoryCount
	idx    int
	done   bool
	err    error
	c      *Client
	ctx    context.Context
}

func (i *batchRetrieveInventoryCountsIterator) setError(err error) bool {
	i.counts = nil
	i.err = err
	return false
}

func (i *batchRetrieveInventoryCountsIterator) Value() *InventoryCount {
	return i.counts[i.idx]
}

func (i *batchRetrieveInventoryCountsIterator) Error() error {
	return i.err
}

func (i *batchRetrieveInventoryCountsIterator) Next() bool {
	i.idx = i.idx + 1
	if i.idx < len(i.counts) {
		return true
	}

	if i.done {
		return false
	}

	body := struct {
		CatalogObjectIDs []string   `json:"catalog_object_ids,omitempty"`
		LocationIDs      []string   `json:"location_ids,omitempty"`
		UpdatedAfter     *time.Time `json:"updated_after,omitempty"`
		Cursor           string     `json:"cursor,omitempty"`
	}{
		CatalogObjectIDs: i.catalogObjectIDs,
		LocationIDs:      i.locationIDs,
		UpdatedAfter:     i.updatedAfter,
		Cursor:           i.cursor,
	}

	jsonBody, err := json.Marshal(&body)
	if err != nil {
		return i.setError(fmt.Errorf("error marshaling request body: %w", err))
	}
	buf := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequest("POST", i.c.endpoint("inventory/batch-retrieve-counts").String(), buf)
	if err != nil {
		return i.setError(fmt.Errorf("error generating new request: %w", err))
	}
	req = req.WithContext(i.ctx)

	resp, err := i.c.httpClient.Do(req)
	if err != nil {
		return i.setError(fmt.Errorf("Error with square http request: %w", err))
	}
	defer resp.Body.Close()

	// If there's an error we still want to try and unmarshal and look for a more
	// descriptive error in the json.  However, if we can't unmarshal the json, we
	// can say that the status code was bad
	var badStatusErr error
	if resp.StatusCode != http.StatusOK {
		badStatusErr = unexpectedCodeError(resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if badStatusErr != nil {
			return i.setError(badStatusErr)
		}
		return i.setError(fmt.Errorf("Error reading response body: %w", err))
	}
	respJson := struct {
		Errors []*Error          `json:"errors"`
		Cursor string            `json:"cursor"`
		Counts []*InventoryCount `json:"counts"`
	}{}
	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		if badStatusErr != nil {
			return i.setError(badStatusErr)
		}
		return i.setError(fmt.Errorf("Error unmarshalling json response: %w", err))
	}
	if len(respJson.Errors) != 0 {
		return i.setError(&ErrorList{respJson.Errors})
	}
	// If we've made it this far with a bad status code, return something
	// about the bad status code
	if badStatusErr != nil {
		return i.setError(badStatusErr)
	}

	if len(respJson.Counts) == 0 {
		return false
	}
	i.counts = respJson.Counts
	i.idx = 0
	if respJson.Cursor == "" {
		i.done = true
	} else {
		i.cursor = respJson.Cursor
	}
	return true
}

func (c *Client) BatchRetrieveInventoryCounts(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) BatchRetrieveInventoryCountsIterator {
	return &batchRetrieveInventoryCountsIterator{
		catalogObjectIDs: catalogObjectIDs,
		locationIDs:      locationIDs,
		updatedAfter:     updatedAfter,
		cursor:           "",
		counts:           nil,
		idx:              -1,
		done:             false,
		err:              nil,
		c:                c,
		ctx:              ctx,
	}
}
