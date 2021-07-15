package mage

import (
	"context"
	"fmt"

	"github.com/magefile/mage/mg"
)

type TerraformInputVars struct {
	Deploy string
}

type TerraformOutputVars struct {
	BackendAddr   string
	MailKey       string
	SquareKey     string
	RecaptchaKey  string
	Auth0Domain   string
	Auth0ClientID string
	Auth0Audience string
}

type TerraformOutputValues struct {
	BackendAddr   string
	MailKey       string
	SquareKey     string
	RecaptchaKey  string
	Auth0Domain   string
	Auth0ClientID string
	Auth0Audience string
}

type TerraformVarType struct {
	Workspace string
	Input     *TerraformInputVars
}

var (
	TerraformVarMap = map[WorkspaceType]*TerraformVarType{
		Testing: &TerraformVarType{
			Workspace: "ws-ngn7vnajVjG8trSu",
			Input: &TerraformInputVars{
				Deploy: "var-oywYeVGkALz94upp",
			},
		},
	}

	terraformOutputs = &TerraformOutputVars{
		BackendAddr:   "backend_addr",
		SquareKey:     "square_key",
		MailKey:       "mail_key",
		RecaptchaKey:  "recaptcha_key",
		Auth0Domain:   "auth0_domain",
		Auth0ClientID: "auth0_client_id",
		Auth0Audience: "auth0_audience",
	}

	terraformOutputValues = &TerraformOutputValues{}
)

func InitTerraformOutputs(ctx context.Context) error {
	mg.Deps(InitTerraformClient, InitWorkspace)

	terraformVars := TerraformVars()

	version, err := TerraformClient().StateVersions.Current(ctx, terraformVars.Workspace)
	if err != nil {
		return fmt.Errorf("error reading state version: %w", err)
	}

	for _, outputVar := range version.Outputs {
		svo, err := TerraformClient().StateVersionOutputs.Read(ctx, outputVar.ID)
		if err != nil {
			return fmt.Errorf("error getting state version output \"%s\": %w", outputVar.ID, err)
		}

		switch svo.Name {
		case terraformOutputs.BackendAddr:
			terraformOutputValues.BackendAddr = svo.Value
		case terraformOutputs.SquareKey:
			terraformOutputValues.SquareKey = svo.Value
		case terraformOutputs.MailKey:
			terraformOutputValues.MailKey = svo.Value
		case terraformOutputs.RecaptchaKey:
			terraformOutputValues.RecaptchaKey = svo.Value
		case terraformOutputs.Auth0Domain:
			terraformOutputValues.Auth0Domain = svo.Value
		case terraformOutputs.Auth0ClientID:
			terraformOutputValues.Auth0ClientID = svo.Value
		case terraformOutputs.Auth0Audience:
			terraformOutputValues.Auth0Audience = svo.Value
		default:
			// Do nothing
		}
	}

	if Workspace() == Local {
		terraformOutputValues.BackendAddr = "http://localhost:8080"
	}

	return nil
}

func TerraformOutputs() *TerraformOutputValues {
	mg.Deps(InitTerraformOutputs)

	return terraformOutputValues
}

func TerraformVars() *TerraformVarType {
	mg.Deps(InitWorkspace)

	workspace := Workspace()
	if Workspace() == Local {
		workspace = Testing
	}

	return TerraformVarMap[workspace]
}
