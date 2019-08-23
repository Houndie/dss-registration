package square

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) BatchRetrieveOrders(ctx context.Context, locationId string, orderIds []string) ([]*Order, error) {
	reqBody := struct {
		OrderIds []string `json:"order_ids"`
	}{
		OrderIds: orderIds,
	}

	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "error marshaling request body")
	}

	bodyBuf := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest("POST", "https://connect.squareup.com/v2/locations/"+locationId+"/orders/batch-retrieve", bodyBuf)
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
		Orders []*Order `json:"orders"`
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
	return respJson.Orders, nil
}
