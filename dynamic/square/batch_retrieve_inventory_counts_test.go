package square

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestBatchRetrieveInventoryCounts(t *testing.T) {
	catalogObjectIds := []string{"id1", "id2"}
	locationIds := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

	cursors := []string{"", "CURSOR", ""}
	callCount := 0

	time1 := time.Unix(1234567, 0)
	time2 := time.Unix(3446678, 0)
	expectedCounts := []*InventoryCount{
		&InventoryCount{
			CatalogObjectId:   catalogObjectIds[0],
			CatalogObjectType: CatalogObjectTypeItemVariation,
			State:             "OH",
			LocationId:        locationIds[0],
			Quantity:          "7",
			CalculatedAt:      &time1,
		},
		&InventoryCount{
			CatalogObjectId:   catalogObjectIds[1],
			CatalogObjectType: CatalogObjectTypeItemVariation,
			State:             "PA",
			LocationId:        locationIds[1],
			Quantity:          "3.4",
			CalculatedAt:      &time2,
		},
	}
	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Authorization") != "Bearer "+apiKey {
					t.Fatalf("Found incorrect authorization header %s, expected %s", r.Header.Get("Authorization"), "Bearer "+apiKey)
				}
				if r.Header.Get("Accept") != "application/json" {
					t.Fatalf("found incorrect accept header %s, expected application/json", r.Header.Get("Accept"))
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Fatalf("found incorrect content-type %s, expected application/json", r.Header.Get("Content-Type"))
				}

				if r.URL.String() != "https://connect.squareup.com/v2/inventory/batch-retrieve-counts" {
					t.Fatalf("Found incorrect request url %s", r.URL.String())
				}

				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("Error reading request body: %v", err)
				}

				jsonRequest := struct {
					CatalogObjectIds []string   `json:"catalog_object_ids"`
					LocationIds      []string   `json:"location_ids"`
					UpdatedAfter     *time.Time `json:"updated_after"`
					Cursor           string     `json:"cursor"`
				}{}
				err = json.Unmarshal(body, &jsonRequest)
				if err != nil {
					t.Fatalf("Error unmarshalling request body: %v", err)
				}

				if len(catalogObjectIds) != len(jsonRequest.CatalogObjectIds) {
					t.Fatalf("wrong number of catalog object ids (found %v, expected %v)", len(jsonRequest.CatalogObjectIds), len(catalogObjectIds))
				}

				for _, controlId := range catalogObjectIds {
					found := false
					for _, testId := range jsonRequest.CatalogObjectIds {
						if testId == controlId {
							found = true
							break
						}
					}
					if !found {
						t.Fatalf("Could not find control catalog id %s", controlId)
					}
				}

				if len(locationIds) != len(jsonRequest.LocationIds) {
					t.Fatalf("wrong number of location ids (found %v, expected %v)", len(jsonRequest.LocationIds), len(locationIds))
				}

				for _, controlId := range locationIds {
					found := false
					for _, testId := range jsonRequest.LocationIds {
						if testId == controlId {
							found = true
							break
						}
					}
					if !found {
						t.Fatalf("Could not find control location id %s", controlId)
					}
				}

				if !updatedAfter.Equal(*jsonRequest.UpdatedAfter) {
					t.Fatalf("Wrong updated after (found %v, expected %v)", jsonRequest.UpdatedAfter, updatedAfter)
				}

				if jsonRequest.Cursor != cursors[callCount] {
					t.Fatalf("incorrect cursor found %s, expected %s", jsonRequest.Cursor, cursors[callCount])
				}

				resp := struct {
					Cursor string
					Counts []*InventoryCount
				}{
					Cursor: cursors[callCount+1],
					Counts: []*InventoryCount{expectedCounts[callCount]},
				}

				jsonResp, err := json.Marshal(&resp)
				if err != nil {
					t.Fatalf("unxpected error marshalling response: %v", err)
				}

				header := http.Header{}
				header.Set("Content-Type", "application/json")

				callCount = callCount + 1
				return &http.Response{
					Status:        http.StatusText(http.StatusOK),
					StatusCode:    http.StatusOK,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					Header:        header,
					Body:          ioutil.NopCloser(bytes.NewReader(jsonResp)),
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
	inventoryCounts := squareClient.BatchRetrieveInventoryCounts(context.Background(), catalogObjectIds, locationIds, &updatedAfter)

	idx := 0
	for inventoryCounts.Next() {
		if inventoryCounts.Value().CatalogObjectId != expectedCounts[idx].CatalogObjectId {
			t.Fatalf("found catalog object id %s, expected %s", inventoryCounts.Value().CatalogObjectId, expectedCounts[idx].CatalogObjectId)
		}
		if inventoryCounts.Value().CatalogObjectType != expectedCounts[idx].CatalogObjectType {
			t.Fatalf("found catalog object type %s, expected %s", inventoryCounts.Value().CatalogObjectType, expectedCounts[idx].CatalogObjectType)
		}
		if inventoryCounts.Value().State != expectedCounts[idx].State {
			t.Fatalf("found state %s, expected %s", inventoryCounts.Value().State, expectedCounts[idx].State)
		}
		if inventoryCounts.Value().LocationId != expectedCounts[idx].LocationId {
			t.Fatalf("found location id %s, expected %s", inventoryCounts.Value().LocationId, expectedCounts[idx].LocationId)
		}
		if inventoryCounts.Value().Quantity != expectedCounts[idx].Quantity {
			t.Fatalf("found quantity %s, expected %s", inventoryCounts.Value().Quantity, expectedCounts[idx].Quantity)
		}
		if !inventoryCounts.Value().CalculatedAt.Equal(*expectedCounts[idx].CalculatedAt) {
			t.Fatalf("found calculated time %s, expected %s", inventoryCounts.Value().CalculatedAt, expectedCounts[idx].CalculatedAt)
		}
		idx = idx + 1
	}

	if inventoryCounts.Error() != nil {
		t.Fatalf("found unexpected error: %v", inventoryCounts.Error())
	}

	if idx != len(expectedCounts) {
		t.Fatalf("found unxepected number of items %v, expected %v", idx, len(expectedCounts))
	}
}

func TestBatchRetrieveInventoryCountsClientError(t *testing.T) {
	catalogObjectIds := []string{"id1", "id2"}
	locationIds := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

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
	inventoryCounts := squareClient.BatchRetrieveInventoryCounts(context.Background(), catalogObjectIds, locationIds, &updatedAfter)

	idx := 0
	for inventoryCounts.Next() {
		idx = idx + 1
	}

	if inventoryCounts.Error() == nil {
		t.Fatalf("expected error, found none")
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}

func TestBatchRetrieveInventoryCountsErrorCode(t *testing.T) {
	catalogObjectIds := []string{"id1", "id2"}
	locationIds := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

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
	inventoryCounts := squareClient.BatchRetrieveInventoryCounts(context.Background(), catalogObjectIds, locationIds, &updatedAfter)

	idx := 0
	for inventoryCounts.Next() {
		idx = idx + 1
	}

	if inventoryCounts.Error() == nil {
		t.Fatalf("expected error, found none")
	}

	uerr, ok := inventoryCounts.Error().(unexpectedCodeError)
	if !ok {
		t.Fatalf("error was not of type unexpectedCodeError")
	}

	if int(uerr) != http.StatusInternalServerError {
		t.Fatalf("error code was not internal server error, found %v", int(uerr))
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}

func TestBatchRetrieveInventoryCountsErrorMessage(t *testing.T) {
	catalogObjectIds := []string{"id1", "id2"}
	locationIds := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

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
	inventoryCounts := squareClient.BatchRetrieveInventoryCounts(context.Background(), catalogObjectIds, locationIds, &updatedAfter)

	idx := 0
	for inventoryCounts.Next() {
		idx = idx + 1
	}

	if inventoryCounts.Error() == nil {
		t.Fatalf("expected error, found none")
	}
	serr, ok := errors.Cause(inventoryCounts.Error()).(*ErrorList)
	if !ok {
		t.Fatalf("error not of type square.Error")
	}
	if len(serr.Errors) != 1 {
		t.Fatalf("found %v errors, expected %v", len(serr.Errors), 1)
	}
	if serr.Errors[0].Category != testError.Category {
		t.Fatalf("found error category %s, expected %s", serr.Errors[0].Category, testError.Category)
	}
	if serr.Errors[0].Code != testError.Code {
		t.Fatalf("found error code %s, expected %s", serr.Errors[0].Code, testError.Code)
	}
	if serr.Errors[0].Detail != testError.Detail {
		t.Fatalf("found error detail %s, expected %s", serr.Errors[0].Detail, testError.Detail)
	}
	if serr.Errors[0].Field != testError.Field {
		t.Fatalf("found error field %s, expected %s", serr.Errors[0].Field, testError.Field)
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}
