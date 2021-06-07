package mage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/magefile/mage/mg"
)

func HealthCheck(ctx context.Context, method, url string) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return fmt.Errorf("error generating new http request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error performing http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("found unexpected status: %s", res.Status)
	}

	body := struct {
		Healthiness string
	}{}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return fmt.Errorf("error decoding json body: %w", err)
	}

	if body.Healthiness != "Healthy" {
		return fmt.Errorf("found healthiness status: %s", body.Healthiness)
	}

	return nil
}

func VersionCheck(ctx context.Context, method, url string) error {
	mg.Deps(InitDeployVersion)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return fmt.Errorf("error generating new http request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error performing http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("found unexpected status: %s", res.Status)
	}

	body := struct {
		Version string
	}{}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return fmt.Errorf("error decoding json body: %w", err)
	}

	if body.Version != DeployVersion() {
		return fmt.Errorf("found version %s, expected %s", body.Version, DeployVersion())
	}

	return nil
}
