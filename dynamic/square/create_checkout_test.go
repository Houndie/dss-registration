package square

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateCheckout(t *testing.T) {
	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*OrderLineItem{
				&OrderLineItem{
					Uid:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &OrderQuantityUnit{
						MeasurementUnit: &MeasurementUnit{
							Type: MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectUrl := "www.mywebsite.com"
	additionalRecipients := []*ChargeRequestAdditionalRecipient{
		&ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

	createdAt := time.Unix(1234567, 0)

	expectedCheckout := &Checkout{
		ID:                         "some checkout id",
		CheckoutPageUrl:            "www.sqaureup.com/gohere",
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectUrl:                redirectUrl,
		Order:                      order.Order,
		CreatedAt:                  &createdAt,
		AdditionalRecipients: []*AdditionalRecipient{
			&AdditionalRecipient{
				LocationID:   additionalRecipients[0].LocationID,
				Description:  additionalRecipients[0].Description,
				AmountMoney:  additionalRecipients[0].AmountMoney,
				ReceivableID: "some receivable id",
			},
		},
	}
	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Method != "POST" {
					t.Fatalf("found non post method %s", r.Method)
				}

				if r.Header.Get("Content-Type") != "application/json" {
					t.Fatalf("found wrong content type header %s", r.Header.Get("Content-Type"))
				}
				if r.Header.Get("Accept") != "application/json" {
					t.Fatalf("found wrong accept header %s", r.Header.Get("Accept"))
				}
				if r.Header.Get("Authorization") != "Bearer "+apiKey {
					t.Fatalf("found wrong authorization header %s", r.Header.Get("Authorization"))
				}

				reqJson := struct {
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

				reqBody, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("error reading request body: %v", err)
				}

				err = json.Unmarshal(reqBody, &reqJson)
				if err != nil {
					t.Fatalf("error unmarshaling request body: %v", err)
				}

				if reqJson.IdempotencyKey != idempotencyKey {
					t.Fatalf("found idepotency key %s, expected %s", reqJson.IdempotencyKey, idempotencyKey)
				}

				if !cmp.Equal(reqJson.Order, order, cmpopts.IgnoreUnexported()) {
					t.Fatalf("found order %s not equal to existing order %s", spew.Sdump(reqJson.Order), spew.Sdump(order))
				}

				if reqJson.AskForShippingAddress != askForShippingAddress {
					t.Fatalf("found ask for shipping param %v, expected %v", reqJson.AskForShippingAddress, askForShippingAddress)
				}

				if reqJson.MerchantSupportEmail != merchantSupportEmail {
					t.Fatalf("found merchant support email %s, expected %s", reqJson.MerchantSupportEmail, merchantSupportEmail)
				}

				if reqJson.PrePopulateBuyerEmail != prePopulateBuyerEmail {
					t.Fatalf("found pre populate buyer email %s, expected %s", reqJson.PrePopulateBuyerEmail, prePopulateBuyerEmail)
				}

				if !cmp.Equal(reqJson.PrePopulateShippingAddress, prePopulateShippingAddress, cmpopts.IgnoreUnexported()) {
					t.Fatalf("found wrong pre populate shipping address %s, expected %s", spew.Sdump(reqJson.PrePopulateShippingAddress), spew.Sdump(prePopulateShippingAddress))
				}

				if reqJson.RedirectUrl != redirectUrl {
					t.Fatalf("found redirect url %s, expected %s", reqJson.RedirectUrl, redirectUrl)
				}

				if !cmp.Equal(reqJson.AdditionalRecipients, additionalRecipients, cmpopts.IgnoreUnexported()) {
					t.Fatalf("found additional recipients %s, expected %s", spew.Sdump(reqJson.AdditionalRecipients), spew.Sdump(additionalRecipients))
				}

				if reqJson.Note != note {
					t.Fatalf("found note %s, expected %s", reqJson.Note, note)
				}

				resJson := struct {
					Checkout *Checkout `json:"checkout"`
				}{
					Checkout: expectedCheckout,
				}

				resBody, err := json.Marshal(&resJson)
				if err != nil {
					t.Fatalf("error marshaling response body: %v", err)
				}

				header := http.Header{}
				header.Set("Content-Type", "application/json")

				return &http.Response{
					Status:        http.StatusText(http.StatusOK),
					StatusCode:    http.StatusOK,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					Header:        header,
					Body:          ioutil.NopCloser(bytes.NewReader(resBody)),
					ContentLength: -1,
					Request:       r,
				}, nil
			},
		},
	}

	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	checkout, err := squareClient.CreateCheckout(context.Background(), locationID, idempotencyKey, order, askForShippingAddress, merchantSupportEmail, prePopulateBuyerEmail, prePopulateShippingAddress, redirectUrl, additionalRecipients, note)
	if err != nil {
		t.Fatalf("found unxpected error from CreateCheckout: %v", err)
	}

	if !cmp.Equal(checkout, expectedCheckout, cmpopts.IgnoreUnexported()) {
		t.Fatalf("found checkout %s, expected %s", spew.Sdump(checkout), spew.Sdump(expectedCheckout))
	}
}

func TestCreateCheckoutClientError(t *testing.T) {
	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*OrderLineItem{
				&OrderLineItem{
					Uid:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &OrderQuantityUnit{
						MeasurementUnit: &MeasurementUnit{
							Type: MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectUrl := "www.mywebsite.com"
	additionalRecipients := []*ChargeRequestAdditionalRecipient{
		&ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("some error")
			},
		},
	}

	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}
	_, err = squareClient.CreateCheckout(context.Background(), locationID, idempotencyKey, order, askForShippingAddress, merchantSupportEmail, prePopulateBuyerEmail, prePopulateShippingAddress, redirectUrl, additionalRecipients, note)

	if err == nil {
		t.Fatal("found no error when client returned one?")
	}
}

func TestCreateCheckoutErrorCode(t *testing.T) {
	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*OrderLineItem{
				&OrderLineItem{
					Uid:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &OrderQuantityUnit{
						MeasurementUnit: &MeasurementUnit{
							Type: MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectUrl := "www.mywebsite.com"
	additionalRecipients := []*ChargeRequestAdditionalRecipient{
		&ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					Status:        http.StatusText(http.StatusInternalServerError),
					StatusCode:    http.StatusInternalServerError,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					ContentLength: 0,
					Body:          ioutil.NopCloser(bytes.NewReader([]byte{})),
					Request:       r,
				}, nil
			},
		},
	}

	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}
	_, err = squareClient.CreateCheckout(context.Background(), locationID, idempotencyKey, order, askForShippingAddress, merchantSupportEmail, prePopulateBuyerEmail, prePopulateShippingAddress, redirectUrl, additionalRecipients, note)

	if err == nil {
		t.Fatal("found no error when client returned one?")
	}

	var uerr unexpectedCodeError
	if !errors.As(err, &uerr) {
		t.Fatalf("error was not of type unexpectedCodeError")
	}

	if int(uerr) != http.StatusInternalServerError {
		t.Fatalf("error code was not internal server error, found %v", int(uerr))
	}
}

func TestCreateCheckoutErrorMessage(t *testing.T) {
	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*OrderLineItem{
				&OrderLineItem{
					Uid:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &OrderQuantityUnit{
						MeasurementUnit: &MeasurementUnit{
							Type: MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectUrl := "www.mywebsite.com"
	additionalRecipients := []*ChargeRequestAdditionalRecipient{
		&ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

	testError := &Error{
		Category: ErrorCategoryApiError,
		Code:     ErrorCodeInternalServerError,
		Detail:   "some detail",
		Field:    "some field",
	}

	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				resp := struct {
					Errors []*Error
				}{
					Errors: []*Error{testError},
				}

				respJson, err := json.Marshal(&resp)
				if err != nil {
					t.Fatalf("error marshaling response body: %v", err)
				}
				return &http.Response{
					Status:        http.StatusText(http.StatusInternalServerError),
					StatusCode:    http.StatusInternalServerError,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					ContentLength: 0,
					Body:          ioutil.NopCloser(bytes.NewReader(respJson)),
					Request:       r,
				}, nil
			},
		},
	}

	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}
	_, err = squareClient.CreateCheckout(context.Background(), locationID, idempotencyKey, order, askForShippingAddress, merchantSupportEmail, prePopulateBuyerEmail, prePopulateShippingAddress, redirectUrl, additionalRecipients, note)

	if err == nil {
		t.Fatal("found no error when client returned one?")
	}

	serr, ok := err.(*ErrorList)
	if !ok {
		t.Fatalf("error was not of type square.ErrorList")
	}

	if len(serr.Errors) != 1 {
		t.Fatalf("found %v errors, expected %v", len(serr.Errors), 1)
	}

	if !cmp.Equal(serr.Errors[0], testError, cmpopts.IgnoreUnexported()) {
		t.Fatalf("errors were not equal, found %s, expected %s", spew.Sdump(serr.Errors[0]), spew.Sdump(testError))
	}
}
