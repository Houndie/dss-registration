package square

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BatchChangeInventoryOption interface {
	isBatchChangeInventoryOption()
}

type IgnoreUnchangedCounts bool

func (*IgnoreUnchangedCounts) isBatchChangeInventoryOption() {}

func (c *Client) BatchChangeInventory(ctx context.Context, idempotencyKey string, changes []*InventoryChange, opts ...BatchChangeInventoryOption) ([]*InventoryCount, error) {
	reqBody := struct {
		IdempotencyKey        string             `json:"idempotency_key,omitempty"`
		Changes               []*InventoryChange `json:"changes,omitempty"`
		IgnoreUnchangedCounts *bool              `json:"ignore_unchanged_counts,omitempty"`
	}{
		IdempotencyKey: idempotencyKey,
		Changes:        changes,
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case *IgnoreUnchangedCounts:
			val := bool(*o)
			reqBody.IgnoreUnchangedCounts = &val
		}
	}

	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	bodyBuf := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest("POST", c.endpoint("inventory/batch-change").String(), bodyBuf)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error with http request: %w", err)
	}
	defer resp.Body.Close()
	var codeErr error
	if resp.StatusCode != http.StatusOK {
		codeErr = unexpectedCodeError(resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if codeErr != nil {
			return nil, codeErr
		}
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	respJson := struct {
		Counts []*InventoryCount `json:"counts"`
		Errors []*Error          `json:"errors"`
	}{}
	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		if codeErr != nil {
			return nil, codeErr
		}
		return nil, fmt.Errorf("error unmarshalling json response: %w", err)
	}
	if len(respJson.Errors) != 0 {
		return nil, &ErrorList{respJson.Errors}
	}
	if codeErr != nil {
		return nil, codeErr
	}
	return respJson.Counts, nil
}
