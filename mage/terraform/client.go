package terraform

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-tfe"
)

type StateVersionOutput struct {
	ID        string
	Name      string
	Sensitive bool
	Type      string
	Value     interface{}
}

type StateVersionOutputs interface {
	Read(ctx context.Context, outputID string) (*StateVersionOutput, error)
}

type stateVersionOutputs struct {
	cfg *tfe.Config
}

func (s stateVersionOutputs) Read(ctx context.Context, outputID string) (*StateVersionOutput, error) {
	var baseURL *url.URL
	var err error
	if s.cfg.Address != "" {
		baseURL, err = url.Parse(s.cfg.Address)
		if err != nil {
			return nil, fmt.Errorf("error parsing baseURL: %w", err)
		}
	} else {
		baseURL, err = url.Parse(tfe.DefaultAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing baseURL: %w", err)
		}
	}

	u, err := baseURL.Parse(fmt.Sprintf("api/v2/state-version-outputs/%s", url.QueryEscape(outputID)))
	if err != nil {
		return nil, fmt.Errorf("error parsing path: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error parsing new request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+s.cfg.Token)
	req.Header.Set("Accept", "application/vnd.api+json")
	req = req.WithContext(ctx)

	res, err := s.cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling state version output read: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("found error code %s", res.Status)
	}

	jsonOutput := struct {
		Data struct {
			ID         string `json:"id"`
			Attributes struct {
				Name      string      `json:"name"`
				Sensitive bool        `json:"sensitive"`
				Type      string      `json:"type"`
				Value     interface{} `json:"value"`
			} `json:"attributes"`
		} `json:"data"`
	}{}

	if err = json.NewDecoder(res.Body).Decode(&jsonOutput); err != nil {
		return nil, fmt.Errorf("error decoding state version output: %w", err)
	}

	return &StateVersionOutput{
		ID:        jsonOutput.Data.ID,
		Name:      jsonOutput.Data.Attributes.Name,
		Sensitive: jsonOutput.Data.Attributes.Sensitive,
		Type:      jsonOutput.Data.Attributes.Type,
		Value:     jsonOutput.Data.Attributes.Value,
	}, nil
}

type Client struct {
	tfe.Client
	StateVersionOutputs StateVersionOutputs
}

func NewClient(config *tfe.Config) (*Client, error) {
	tfeClient, err := tfe.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error creating internal client: %w", err)
	}
	return &Client{
		Client: *tfeClient,
		StateVersionOutputs: stateVersionOutputs{
			cfg: config,
		},
	}, nil
}
