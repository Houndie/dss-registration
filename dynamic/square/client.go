package square

import "net/http"

type Client struct {
	httpClient *http.Client
}

func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		return &Client{
			httpClient: &http.Client{
				Transport: &middleware{
					apiKey: apiKey,
					wrap:   http.DefaultTransport,
				},
			},
		}
	}

	var transport http.RoundTripper
	if httpClient.Transport == nil {
		transport = &middleware{
			apiKey: apiKey,
			wrap:   http.DefaultTransport,
		}
	} else {
		transport = &middleware{
			apiKey: apiKey,
			wrap:   httpClient.Transport,
		}
	}
	return &Client{
		httpClient: &http.Client{
			Transport:     transport,
			CheckRedirect: httpClient.CheckRedirect,
			Jar:           httpClient.Jar,
			Timeout:       httpClient.Timeout,
		},
	}
}

type middleware struct {
	apiKey string
	wrap   http.RoundTripper
}

func (m middleware) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bearer "+m.apiKey)
	r.Header.Add("Accept", "application/json")
	if r.Method == "POST" {
		r.Header.Add("Content-Type", "application/json")
	}
	return m.wrap.RoundTrip(r)
}
