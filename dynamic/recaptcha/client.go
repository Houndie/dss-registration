package recaptcha

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	secretKey  string
	verifyURL  *url.URL
}

func NewClient(httpClient *http.Client, secretKey string) (*Client, error) {
	u, err := url.Parse("https://www.google.com/recaptcha/api/siteverify")
	if err != nil {
		return nil, fmt.Errorf("error parsing verify url: %w", err)
	}
	return &Client{
		httpClient: httpClient,
		secretKey:  secretKey,
		verifyURL:  u,
	}, nil
}

func (c *Client) VerifySite(ctx context.Context, recaptchaResponse string) (bool, error) {
	u := &url.URL{
		Scheme:      c.verifyURL.Scheme,
		Opaque:      c.verifyURL.Opaque,
		User:        c.verifyURL.User,
		Host:        c.verifyURL.Host,
		Path:        c.verifyURL.Path,
		RawPath:     c.verifyURL.RawPath,
		ForceQuery:  c.verifyURL.ForceQuery,
		Fragment:    c.verifyURL.Fragment,
		RawFragment: c.verifyURL.RawFragment,
	}
	q := c.verifyURL.Query()
	q.Add("response", recaptchaResponse)
	q.Add("secret", c.secretKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return false, fmt.Errorf("error creating request: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("error with http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response body: %w", err)
	}

	respJson := struct {
		Success bool `json:"success"`
	}{}
	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		return false, fmt.Errorf("error unmarshalling json response: %w", err)
	}
	return respJson.Success, nil

}
