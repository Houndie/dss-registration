package square

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type BatchRetrieveInventoryCountsIterator interface {
	Value() *InventoryCount
	Error() error
	Next() bool
}

type batchRetrieveInventoryCountsIterator struct {
	catalogObjectIds []string
	locationIds      []string
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

	url, err := url.Parse("https://connect.squareup.com/v2/inventory/batch-retrieve-counts")
	if err != nil {
		return i.setError(errors.Wrap(err, "Error parsing url"))
	}
	body := struct {
		CatalogObjectIds []string   `json:"catalog_object_ids,omitempty"`
		LocationIds      []string   `json:"location_ids,omitempty"`
		UpdatedAfter     *time.Time `json:"updated_after,omitempty"`
		Cursor           string     `json:"cursor,omitempty"`
	}{
		CatalogObjectIds: i.catalogObjectIds,
		LocationIds:      i.locationIds,
		UpdatedAfter:     i.updatedAfter,
		Cursor:           i.cursor,
	}

	jsonBody, err := json.Marshal(&body)
	if err != nil {
		return i.setError(errors.Wrap(err, "error marshaling request body"))
	}
	buf := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequest("POST", url.String(), buf)
	if err != nil {
		return i.setError(errors.Wrap(err, "error generating new request"))
	}
	req = req.WithContext(i.ctx)

	resp, err := i.c.httpClient.Do(req)
	if err != nil {
		return i.setError(errors.Wrap(err, "Error with square http request"))
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
		return i.setError(errors.Wrap(err, "Error reading response body"))
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
		return i.setError(errors.Wrap(err, "Error unmarshalling json response"))
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

func (c *Client) BatchRetrieveInventoryCounts(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) BatchRetrieveInventoryCountsIterator {
	return &batchRetrieveInventoryCountsIterator{
		catalogObjectIds: catalogObjectIds,
		locationIds:      locationIds,
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
