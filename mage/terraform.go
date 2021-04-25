package mage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

func TerraformWorkspaceID(ctx context.Context, client *http.Client, workspaceName string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "/organizations/daytonswingsmackdown/workspaces", nil)
	if err != nil {
		return "", fmt.Errorf("error creating workspaces request: %w", err)
	}

	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error listing worksapces: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("found unexpected error code %d when listing workspaces", res.StatusCode)
	}

	workspaces := struct {
		Data []struct {
			ID         string `json:"id"`
			Attributes struct {
				Name string `json:"name"`
			} `json:"attributes"`
		} `json:"data"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&workspaces); err != nil {
		return "", fmt.Errorf("error decoding workspaces list response: %w", err)
	}

	var workspaceID string
	for _, w := range workspaces.Data {
		if w.Attributes.Name == workspaceName {
			workspaceID = w.ID
		}
	}

	if workspaceID == "" {
		return "", fmt.Errorf("workspace with name %s not found", workspaceName)
	}

	return workspaceID, nil
}

const deployVersionKey = "deploy_version"

func TerraformDeployVersionID(ctx context.Context, client *http.Client, workspaceID string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/workspaces/%s/vars", workspaceID), nil)
	if err != nil {
		return "", fmt.Errorf("error creating workspace vars request: %w", err)
	}

	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error listing workspace vars: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("found unexpected error code %d when listing workspace vars", res.StatusCode)
	}

	workspaceVars := struct {
		Data []struct {
			ID         string `json:"id"`
			Attributes struct {
				Key string `json:"key"`
			} `json:"attributes"`
		} `json:"data"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&workspaceVars); err != nil {
		return "", fmt.Errorf("error decoding workspace vars list response: %w", err)
	}

	var varID string
	for _, v := range workspaceVars.Data {
		if v.Attributes.Key == deployVersionKey {
			varID = v.ID
		}
	}

	if varID == "" {
		return "", fmt.Errorf("workspace var with name %s not found", deployVersionKey)
	}

	return varID, nil
}

type updateVarData struct {
	Type       string              `json:"type"`
	ID         string              `json:"id"`
	Attributes updateVarAttributes `json:"attributes"`
}

type updateVarAttributes struct {
	Value string `json:"value"`
}

func SetTerraformDeployVersionID(ctx context.Context, client *http.Client, workspaceID, varID, deployVersion string) error {
	body, err := json.Marshal(struct {
		Data updateVarData `json:"data"`
	}{
		Data: updateVarData{
			Type: "vars",
			ID:   varID,
			Attributes: updateVarAttributes{
				Value: deployVersion,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("error marshaling set var request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/workspaces/%s/vars/%s", workspaceID, varID), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating set workspace var request: %w", err)
	}

	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error setting workspace var: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("found unexpected error code %d when setting workspace var", res.StatusCode)
	}

	return nil
}

type TerraformTransport struct {
	ApiKey string
}

func (t *TerraformTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("authorization", "Bearer "+t.ApiKey)

	u, err := url.Parse("https://" + path.Join("app.terraform.io/api/v2/", r.URL.String()))
	if err != nil {
		return nil, fmt.Errorf("error parsing terraform url: %w", err)
	}

	r.URL = u

	if r.Method != http.MethodGet {
		r.Header.Add("content-type", "application/vnd.api+json")
	}

	return http.DefaultTransport.RoundTrip(r)
}
