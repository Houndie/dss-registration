package google

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockRoundTripper struct {
	RoundTripFunc func(r *http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(r)
}

func TestUserinfo(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserinfoEndpoint := "https://getuserinfo"
	testUserId := "12345"
	discoveryDocumentCount := 0

	client := &http.Client{
		Transport: &mockRoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				switch r.URL.String() {
				case discoveryDocumentURI:
					discoveryDocumentCount++

					body := struct {
						UserinfoEndpoint string `json:"userinfo_endpoint"`
					}{
						UserinfoEndpoint: testUserinfoEndpoint,
					}

					bodyBytes, err := json.Marshal(&body)
					if err != nil {
						t.Fatalf("error marshaling discovery document response body: %v", err)
					}
					header := http.Header{}
					header.Add("Cache-Control", "max-age=5000")
					return &http.Response{
						Status:        http.StatusText(http.StatusOK),
						StatusCode:    http.StatusOK,
						Header:        header,
						Proto:         "HTTP/1.0",
						ProtoMajor:    1,
						ProtoMinor:    0,
						ContentLength: 0,
						Body:          ioutil.NopCloser(bytes.NewReader(bodyBytes)),
						Request:       r,
					}, nil
				case testUserinfoEndpoint:
					if r.Header.Get("Authorization") != "Bearer "+testAccessToken {
						t.Fatalf("found incorrect authorization header %s", r.Header.Get("Authorization"))
					}

					body := struct {
						Sub string `json:"sub"`
					}{
						Sub: testUserId,
					}

					bodyBytes, err := json.Marshal(&body)
					if err != nil {
						t.Fatalf("error marshaling userinfo response body: %v", err)
					}
					return &http.Response{
						Status:        http.StatusText(http.StatusOK),
						StatusCode:    http.StatusOK,
						Proto:         "HTTP/1.0",
						ProtoMajor:    1,
						ProtoMinor:    0,
						ContentLength: 0,
						Body:          ioutil.NopCloser(bytes.NewReader(bodyBytes)),
						Request:       r,
					}, nil
				default:
					t.Fatalf("found unknown url %s", r.URL.String())
				}
				return nil, nil
			},
		},
	}

	userinfo, err := NewAuthorizer(client).Userinfo(context.Background(), testAccessToken)
	if err != nil {
		t.Fatalf("found unexpected error when fetching user info: %v", err)
	}

	if userinfo.UserId != testUserId {
		t.Fatalf("found user id %s, expected %s", userinfo.UserId, testUserId)
	}

	if discoveryDocumentCount != 1 {
		t.Fatalf("found more than one call to discovery document (%d)", discoveryDocumentCount)
	}
}
