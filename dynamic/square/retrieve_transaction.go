package square

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) RetrieveTransaction(ctx context.Context, locationId, transactionId string) (*Transaction, error) {
	req, err := http.NewRequest("GET", "https://connect.squareup.com/v2/locations/"+locationId+"/transactions/"+transactionId, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating request")
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Error with http request")
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
		return nil, errors.Wrap(err, "Error reading response body")
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
		return nil, errors.Wrap(err, "Error unmarshalling json response")
	}
	if len(respJson.Errors) != 0 {
		return nil, &ErrorList{respJson.Errors}
	}
	if codeErr != nil {
		return nil, codeErr
	}
	return respJson.Transaction, nil
}
