package square

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) UpdateOrder(ctx context.Context, locationID, orderID string, order *Order, fieldsToClear []string, idempotencyKey string) (*Order, error) {
	reqBody := struct {
		Order          *Order   `json:"order"`
		FieldsToClear  []string `json:"fields_to_clear"`
		IdempotencyKey string   `json:"idempotency_key"`
	}{
		Order:          order,
		FieldsToClear:  fieldsToClear,
		IdempotencyKey: idempotencyKey,
	}

	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "error marshaling request body")
	}

	bodyBuf := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest("PUT", "http://connect.squareup.com/v2/locations/"+locationID+"/orders/"+orderID, bodyBuf)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request")
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error with http request")
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
		return nil, errors.Wrap(err, "error reading response body")
	}

	respJson := struct {
		Order  *Order   `json:"order"`
		Errors []*Error `json:"errors"`
	}{}
	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		if codeErr != nil {
			return nil, codeErr
		}
		return nil, errors.Wrap(err, "error unmarshalling json response")
	}
	if len(respJson.Errors) != 0 {
		return nil, &ErrorList{respJson.Errors}
	}
	if codeErr != nil {
		return nil, codeErr
	}
	return respJson.Order, nil
}
