package square

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) RetrieveTransaction(ctx context.Context, locationID, transactionID string) (*Transaction, error) {
	req, err := http.NewRequest("GET", "https://connect.squareup.com/v2/locations/"+locationID+"/transactions/"+transactionID, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error with http request: %w", err)
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
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	respJson := struct {
		Errors      []*Error     `json:"errors"`
		Transaction *Transaction `json:"transaction"`
	}{}
	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		if codeErr != nil {
			return nil, codeErr
		}
		return nil, fmt.Errorf("Error unmarshalling json response: %w", err)
	}
	if len(respJson.Errors) != 0 {
		return nil, &ErrorList{respJson.Errors}
	}
	if codeErr != nil {
		return nil, codeErr
	}
	return respJson.Transaction, nil
}
