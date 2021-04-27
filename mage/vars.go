package mage

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-tfe"
)

var (
	workspace       string
	deployVersion   string
	terraformClient *tfe.Client
	herokuAPIKey    string
)

func Workspace() string {
	if workspace == "" {
		panic("workspace not initialized")
	}

	return workspace
}

func InitWorkspace() error {
	var ok bool
	workspace, ok = os.LookupEnv("WORKSPACE")
	if !ok {
		return errors.New("environment variable WORKSPACE must not be empty")
	}

	return nil
}

func DeployVersion() string {
	if deployVersion == "" {
		panic("deployVersion not initialized")
	}

	return deployVersion
}

func InitDeployVersion() error {
	var ok bool
	deployVersion, ok = os.LookupEnv("DEPLOY_VERSION")
	if !ok {
		return errors.New("environment variable DEPLOY_VERSION must not be empty")
	}

	return nil
}

func TerraformClient() *tfe.Client {
	if terraformClient == nil {
		panic("terraformClient not initialized")
	}

	return terraformClient
}

func InitTerraformClient() error {
	terraformAPIKey, ok := os.LookupEnv("TERRAFORM_API_KEY")
	if !ok {
		return errors.New("environment variable TERRAFORM_API_KEY must not be empty")
	}

	var err error
	terraformClient, err = tfe.NewClient(&tfe.Config{
		Token: terraformAPIKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		return fmt.Errorf("error creating terraform client: %w", err)
	}

	return nil
}

func HerokuAPIKey() string {
	if herokuAPIKey == "" {
		panic("herokuAPIKey not initialize")
	}

	return herokuAPIKey
}

func InitHerokuAPIKey() error {
	var ok bool
	herokuAPIKey, ok = os.LookupEnv("HEROKU_API_KEY")
	if !ok {
		return errors.New("environment variable HEROKU_API_KEY must not be empty")
	}

	return nil
}
