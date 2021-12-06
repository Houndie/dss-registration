package sendinblue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const endpoint = "https://api.sendinblue.com/v3"

type Client struct {
	httpClient   *http.Client
	apiKey       string
	endpointBase *url.URL
}

type middleware struct {
	apiKey string
	wrap   http.RoundTripper
}

func (m middleware) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("api-key", m.apiKey)
	r.Header.Add("Accept", "application/json")
	if r.Method == "POST" {
		r.Header.Add("Content-Type", "application/json")
	}
	return m.wrap.RoundTrip(r)
}

func NewClient(apiKey string, httpClient *http.Client) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint url: %w", err)
	}
	if httpClient == nil {
		return &Client{
			endpointBase: u,
			httpClient: &http.Client{
				Transport: &middleware{
					apiKey: apiKey,
					wrap:   http.DefaultTransport,
				},
			},
		}, nil
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
		endpointBase: u,
		httpClient: &http.Client{
			Transport:     transport,
			CheckRedirect: httpClient.CheckRedirect,
			Jar:           httpClient.Jar,
			Timeout:       httpClient.Timeout,
		},
	}, nil
}

func (c *Client) endpoint(e string) *url.URL {
	u := &url.URL{
		Scheme: c.endpointBase.Scheme,
		User:   c.endpointBase.User,
		Host:   c.endpointBase.Host,
		Path:   path.Join(c.endpointBase.Path, e),
	}
	return u
}

type Sender struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	ID    int64  `json:"id,omitempty"`
}

type EmailPerson struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Attachment struct {
	URL     string `json:"url,omitempty"`
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
}

type MessageVersion struct {
	To      []*EmailPerson    `json:"to,omitempty"`
	Params  map[string]string `json:"params,omitempty"`
	CC      []*EmailPerson    `json:"cc,omitempty"`
	BCC     []*EmailPerson    `json:"bcc,omitempty"`
	Subject string            `json:"subject,omitempty"`
	ReplyTo *EmailPerson      `json:"replyTo,omitempty"`
}

type SMTPEmailParams struct {
	Sender          *Sender           `json:"sender,omitempty"`
	To              []*EmailPerson    `json:"to,omitempty"`
	CC              []*EmailPerson    `json:"cc,omitempty"`
	BCC             []*EmailPerson    `json:"bcc,omitempty"`
	HtmlContent     string            `json:"htmlContent,omitempty"`
	TextContent     string            `json:"textContent,omitempty"`
	Subject         string            `json:"subject,omitempty"`
	ReplyTo         *EmailPerson      `json:"replyTo,omitempty"`
	Attachment      []*Attachment     `json:"attachment,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	TemplateID      int64             `json:"templateId,omitempty"`
	Params          interface{}       `json:"params,omitempty"`
	MessageVersions []*MessageVersion `json:"messageVersions,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
}

func (c *Client) SendSMTPEmail(ctx context.Context, params *SMTPEmailParams) (string, error) {
	jsonBody, err := json.Marshal(params)
	if err != nil {
		return "", fmt.Errorf("error mashaling request body: %w", err)
	}
	bodyBuf := bytes.NewBuffer(jsonBody)
	fmt.Println(bodyBuf.String())

	req, err := http.NewRequest("POST", c.endpoint("smtp/email").String(), bodyBuf)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error with http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bytes))
		return "", fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	respJson := struct {
		MessageID string `json:"messageId"`
	}{}
	err = json.Unmarshal(bytes, &respJson)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling json response: %w", err)
	}
	return respJson.MessageID, nil
}
