package square

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) CreateCheckout(ctx context.Context, locationId, idempotencyKey string, order *CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail string, prePopulateBuyerEmail string, prePopulateShippingAddress *Address, redirectUrl string, additionalRecipients []*ChargeRequestAdditionalRecipient, note string) (*Checkout, error) {
	body := struct {
		IdempotencyKey             string                              `json:"idempotency_key,omitempty"`
		Order                      *CreateOrderRequest                 `json:"order,omitempty"`
		AskForShippingAddress      bool                                `json:"ask_for_shipping_address,omitempty"`
		MerchantSupportEmail       string                              `json:"merchant_support_email,omitempty"`
		PrePopulateBuyerEmail      string                              `json:"pre_populate_buyer_email,omitempty"`
		PrePopulateShippingAddress *Address                            `json:"pre_populate_shipping_address,omitempty"`
		RedirectUrl                string                              `json:"redirect_url,omitempty"`
		AdditionalRecipients       []*ChargeRequestAdditionalRecipient `json:"additional_recipients,omitempty"`
		Note                       string                              `json:"note,omitempty"`
	}{
		IdempotencyKey:             idempotencyKey,
		Order:                      order,
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectUrl:                redirectUrl,
		AdditionalRecipients:       additionalRecipients,
		Note:                       note,
	}
	jsonBody, err := json.Marshal(&body)
	if err != nil {
		return nil, errors.Wrap(err, "Error mashaling request body")
	}

	bodyBuf := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequest("POST", "https://connect.squareup.com/v2/locations/"+locationId+"/checkouts", bodyBuf)
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
		Errors   []*Error  `json:"errors"`
		Checkout *Checkout `json:"checkout"`
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
	return respJson.Checkout, nil
}
