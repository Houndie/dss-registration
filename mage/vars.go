package mage

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-tfe"
	"github.com/magefile/mage/mg"
	"github.com/moby/moby/client"
)

var (
	workspace       WorkspaceType
	deployVersion   string
	terraformClient *tfe.Client
	dockerClient    *client.Client
	herokuAPIKey    string
	migrationURL    string
	dockerCache     string
)

type WorkspaceType string

const (
	Local      WorkspaceType = "local"
	Testing    WorkspaceType = "testing"
	Production WorkspaceType = "production"
)

func Workspace() WorkspaceType {
	if workspace == "" {
		panic("workspace not initialized")
	}

	return workspace
}

func InitWorkspace() error {
	workspaceStr, ok := os.LookupEnv("WORKSPACE")
	if !ok {
		return errors.New("environment variable WORKSPACE must not be empty")
	}
	switch workspaceStr {
	case "testing":
		workspace = Testing
	case "production":
		workspace = Production
	case "local":
		workspace = Local
	default:
		return fmt.Errorf("unknown workspace found: %s", workspaceStr)
	}

	if mg.Verbose() {
		fmt.Fprintf(os.Stderr, "using workspace: %s\n", workspace)
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

	if mg.Verbose() {
		fmt.Fprintf(os.Stderr, "using deploy version: %s\n", deployVersion)
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

func InitMigrationURL() error {
	var ok bool
	migrationURL, ok = os.LookupEnv("MIGRATION_URL")
	if !ok {
		return errors.New("environment variable MIGRATION_URL must not be empty")
	}

	if mg.Verbose() {
		fmt.Fprintf(os.Stderr, "using migration url: %s\n", migrationURL)
	}

	return nil
}

func MigrationURL() string {
	if migrationURL == "" {
		panic("migrationURL not initialized")
	}

	return migrationURL
}

func InitDockerClient() error {
	var err error
	dockerClient, err = client.NewClientWithOpts(client.WithTimeout(10*time.Minute), client.FromEnv)
	if err != nil {
		return fmt.Errorf("error creating new client: %w", err)
	}

	return nil
}

func DockerClient() *client.Client {
	if dockerClient == nil {
		panic("docker client not initialized")
	}

	return dockerClient
}

func InitDockerCache() error {
	var ok bool
	dockerCache, ok = os.LookupEnv("DOCKER_CACHE")
	if !ok {
		return errors.New("environment variable DOCKER_CACHE must not be empty")
	}

	if mg.Verbose() {
		fmt.Fprintf(os.Stderr, "using docker cache: %s\n", dockerCache)
	}

	return nil
}

func DockerCache() string {
	if dockerCache == "" {
		panic("dockerCache not initialized")
	}

	return dockerCache
}
