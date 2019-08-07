package square

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestListCatalog(t *testing.T) {
	apiKey := "some api key"
	types := []CatalogObjectType{
		CatalogObjectTypeTax,
		CatalogObjectTypeModifier,
	}

	updatedAt := time.Unix(1235634, 0)
	updatedAt2 := time.Unix(2363094, 0)
	expectedObjects := []*CatalogObject{
		&CatalogObject{
			Id:                    "some id",
			UpdatedAt:             &updatedAt,
			Version:               7,
			IsDeleted:             true,
			CatalogV1Ids:          nil,
			PresentAtAllLocations: true,
			PresentAtLocationIds:  nil,
			ImageId:               "some image id",
			CatalogObjectType: &CatalogTax{
				Name:                   "tax",
				CalculationPhase:       "phase",
				InclusionType:          "inclusion",
				Percentage:             "6",
				AppliesToCustomAmounts: true,
				Enabled:                true,
			},
		},
		&CatalogObject{
			Id:                    "some other id",
			UpdatedAt:             &updatedAt2,
			Version:               2,
			IsDeleted:             false,
			CatalogV1Ids:          nil,
			PresentAtAllLocations: false,
			PresentAtLocationIds:  []string{"location 1", "location 2"},
			ImageId:               "some other image id",
			CatalogObjectType: &CatalogModifier{
				Name: "modifier",
				PriceMoney: &Money{
					Amount:   3,
					Currency: "pesos",
				},
			},
		},
	}

	cursors := []string{"", "cursor", ""}

	callCount := 0
	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Authorization") != "Bearer "+apiKey {
					t.Fatalf("Found incorrect authorization header %s", r.Header.Get("Authorization"))
				}

				if r.Header.Get("Accept") != "application/json" {
					t.Fatalf("Found incorrect accept header %s", r.Header.Get("Accept"))
				}

				urlParams := struct {
					Types  string `schema:"types"`
					Cursor string `schema:"cursor"`
				}{}

				err := decoder.Decode(&urlParams, r.URL.Query())
				if err != nil {
					t.Fatalf("unexpected error when decoding url params")
				}

				for _, controlType := range types {
					found := false
					for _, testType := range strings.Split(urlParams.Types, ",") {
						if controlType == CatalogObjectType(testType) {
							found = true
							break
						}
					}
					if !found {
						t.Fatalf("could not find expected type %s in params", controlType)
					}
				}

				if cursors[callCount] != urlParams.Cursor {
					t.Fatalf("found unexpected cursor %s, expected %s", cursors[callCount], urlParams.Cursor)
				}

				body, err := json.Marshal(&struct {
					Cursor  string           `json:"cursor,omitempty"`
					Objects []*CatalogObject `json:"objects,omitempty"`
				}{
					Cursor:  cursors[callCount+1],
					Objects: []*CatalogObject{expectedObjects[callCount]},
				})

				if err != nil {
					t.Fatalf("found error while marshaling json response: %v", err)
				}

				callCount = callCount + 1

				return &http.Response{
					Status:        http.StatusText(http.StatusOK),
					StatusCode:    http.StatusOK,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					ContentLength: 0,
					Body:          ioutil.NopCloser(bytes.NewReader(body)),
					Request:       r,
				}, nil
			},
		},
	}
	catalogObjects := NewClient(apiKey, client).ListCatalog(context.Background(), types)

	idx := 0
	for catalogObjects.Next() {
		if !reflect.DeepEqual(catalogObjects.Value(), expectedObjects[idx]) {
			t.Fatalf("found unexpected catalog item %#v, expected %#v", catalogObjects.Value(), expectedObjects[idx])
		}
		idx = idx + 1
	}

	if catalogObjects.Error() != nil {
		t.Fatalf("found unexpected error: %v", catalogObjects.Error())
	}

	if callCount != 2 {
		t.Fatalf("found %d http calls, expected 2", callCount)
	}

	if idx != 2 {
		t.Fatalf("found %d response items, expected 2", idx)
	}
}

func TestListCatalogClientError(t *testing.T) {
	apiKey := "some api key"
	types := []CatalogObjectType{
		CatalogObjectTypeTax,
		CatalogObjectTypeModifier,
	}

	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("some error")
			},
		},
	}
	catalogObjects := NewClient(apiKey, client).ListCatalog(context.Background(), types)

	idx := 0
	for catalogObjects.Next() {
		idx = idx + 1
	}

	if catalogObjects.Error() == nil {
		t.Fatal("Expected error, found none")
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}

func TestListCatalogHttpError(t *testing.T) {
	apiKey := "some api key"
	types := []CatalogObjectType{
		CatalogObjectTypeTax,
		CatalogObjectTypeModifier,
	}

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
	catalogObjects := NewClient(apiKey, client).ListCatalog(context.Background(), types)

	idx := 0
	for catalogObjects.Next() {
		idx = idx + 1
	}

	if catalogObjects.Error() == nil {
		t.Fatal("Expected error, found none")
	}

	uerr, ok := catalogObjects.Error().(unexpectedCodeError)
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

func TestListCatalogErrorMessage(t *testing.T) {
	apiKey := "some api key"
	types := []CatalogObjectType{
		CatalogObjectTypeTax,
		CatalogObjectTypeModifier,
	}

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
	catalogObjects := NewClient(apiKey, client).ListCatalog(context.Background(), types)

	idx := 0
	for catalogObjects.Next() {
		idx = idx + 1
	}

	if catalogObjects.Error() == nil {
		t.Fatal("Expected error, found none")
	}

	serr, ok := errors.Cause(catalogObjects.Error()).(*ErrorList)
	if !ok {
		t.Fatalf("error not of type square.ErrorList")
	}

	if !reflect.DeepEqual(serr.Errors[0], testError) {
		t.Fatalf("expected error %#v, found %#v", serr.Errors[0], testError)
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}
