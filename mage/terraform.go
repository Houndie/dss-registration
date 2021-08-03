package mage

import (
	"context"
	"errors"
	"fmt"

	"github.com/magefile/mage/mg"
)

type TerraformInputVars struct {
	Deploy string
}

type TerraformOutputVars struct {
	BackendAddr                string
	BackendConfigVars          string
	BackendSensitiveConfigVars string
	FrontendConfigVars         string
}

type TerraformOutputValues struct {
	BackendAddr                string
	BackendConfigVars          map[string]string
	BackendSensitiveConfigVars map[string]string
	FrontendConfigVars         map[string]string
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
		BackendAddr:                "backend_addr",
		BackendConfigVars:          "backend_config_vars",
		BackendSensitiveConfigVars: "backend_sensitive_config_vars",
		FrontendConfigVars:         "frontend_config_vars",
	}

	terraformOutputValues = &TerraformOutputValues{}
)

func parseConfigVars(input interface{}) (map[string]string, error) {
	vars, ok := input.(map[string]interface{})
	if !ok {
		return nil, errors.New("config vars are not map")
	}

	output := map[string]string{}
	for key, value := range vars {
		stringValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("config var value for %s is not string", key)
		}

		output[key] = stringValue
	}

	return output, nil
}

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
			ok := true
			terraformOutputValues.BackendAddr, ok = svo.Value.(string)
			if !ok {
				return errors.New("error parsing backend addr as string")
			}
		case terraformOutputs.BackendConfigVars:
			terraformOutputValues.BackendConfigVars, err = parseConfigVars(svo.Value)
			if err != nil {
				return fmt.Errorf("error parsing backend config vars: %w", err)
			}
		case terraformOutputs.BackendSensitiveConfigVars:
			terraformOutputValues.BackendSensitiveConfigVars, err = parseConfigVars(svo.Value)
			if err != nil {
				return fmt.Errorf("error parsing backend sensitive config vars: %w", err)
			}
		case terraformOutputs.FrontendConfigVars:
			terraformOutputValues.FrontendConfigVars, err = parseConfigVars(svo.Value)
			if err != nil {
				return fmt.Errorf("error parsing frontend config vars: %w", err)
			}
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
