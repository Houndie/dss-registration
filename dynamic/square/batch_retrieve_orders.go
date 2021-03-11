package square

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) BatchRetrieveOrders(ctx context.Context, locationID string, orderIDs []string) ([]*Order, error) {
	reqBody := struct {
		OrderIDs []string `json:"order_ids"`
	}{
		OrderIDs: orderIDs,
	}

	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	bodyBuf := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest("POST", c.endpoint("locations/"+locationID+"/orders/batch-retrieve").String(), bodyBuf)
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
		Orders []*Order `json:"orders"`
		Errors []*Error `json:"errors"`
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
	return respJson.Orders, nil
}
