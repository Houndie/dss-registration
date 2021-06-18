package auth0

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/sirupsen/logrus"
)

type mockRoundTripper struct {
	RoundTripFunc func(r *http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(r)
}

func TestUserinfo(t *testing.T) {
	testJWKSEndpoint := "https://endpoint/jwks"
	testUserID := "12345"
	discoveryDocumentCount := 0
	jwksCount := 0
	myEndpoint := "https://endpoint"
	discoveryDocumentURI := "https://endpoint/.well-known/openid-configuration"
	keyID := "keyid"
	permission := authorizer.ListDiscountsPermission

	signingKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("error generating rsa key: %v", err)
	}
	jwkPrivateKey, err := jwk.New(signingKey)
	if err != nil {
		t.Fatalf("error converting key to jwkType: %v", err)
	}
	err = jwkPrivateKey.Set(jwk.KeyIDKey, keyID)
	if err != nil {
		t.Fatalf("error setting kid: %v", err)
	}

	jwkPublicKey, err := jwk.New(signingKey.PublicKey)
	if err != nil {
		t.Fatalf("error converting public key to jwkType: %v", err)
	}
	err = jwkPublicKey.Set(jwk.KeyIDKey, keyID)
	if err != nil {
		t.Fatalf("error setting public kid: %v", err)
	}

	err = jwkPublicKey.Set(jwk.AlgorithmKey, jwa.RS256)
	if err != nil {
		t.Fatalf("error setting public kid: %v", err)
	}
	keySet := jwk.NewSet()
	keySet.Add(jwkPublicKey)

	token := jwt.New()
	err = token.Set(jwt.SubjectKey, testUserID)
	if err != nil {
		t.Fatalf("error adding subject to jwt: %v", err)
	}
	err = token.Set("permissions", []authorizer.Permission{permission})
	if err != nil {
		t.Fatalf("error adding permissions to jwt: %v", err)
	}
	tokenBytes, err := jwt.Sign(token, jwa.RS256, jwkPrivateKey)
	if err != nil {
		t.Fatalf("error signing jwt: %v", err)
	}

	client := &http.Client{
		Transport: &mockRoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				switch r.URL.String() {
				case discoveryDocumentURI:
					discoveryDocumentCount++

					body := struct {
						JWKSURI string `json:"jwks_uri"`
					}{
						JWKSURI: testJWKSEndpoint,
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
				case testJWKSEndpoint:
					jwksCount++
					bodyBytes, err := json.Marshal(&keySet)
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

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})
	authorizer, err := NewAuthorizer(myEndpoint, client, logger)
	if err != nil {
		t.Fatalf("error creating authorizer: %v", err)
	}
	userinfo, err := authorizer.GetUserinfo(context.Background(), string(tokenBytes))
	if err != nil {
		t.Fatalf("found unexpected error when fetching user info: %v", err)
	}

	if userinfo.UserID() != testUserID {
		t.Fatalf("found user id %s, expected %s", userinfo.UserID(), testUserID)
	}

	if !userinfo.IsAllowed(permission) {
		t.Fatalf("userinfo says not allowed for permission")
	}

	if discoveryDocumentCount != 1 {
		t.Fatalf("found more than one call to discovery document (%d)", discoveryDocumentCount)
	}

	if jwksCount != 1 {
		t.Fatalf("found more than one call to jwks endpoint (%d)", discoveryDocumentCount)
	}
}
