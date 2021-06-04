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
	BackendAddr string
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
		BackendAddr: "backend_addr",
	}

	backendAddr        string
	initializedOutputs map[string]struct{}
)

func initOutputVars(ctx context.Context, stopAfter string) error {
	mg.Deps(InitTerraformClient)

	terraformVars, err := TerraformVars()
	if err != nil {
		return err
	}

	version, err := TerraformClient().StateVersions.Current(ctx, terraformVars.Workspace)
	if err != nil {
		return fmt.Errorf("error reading state version: %w", err)
	}

	for _, outputVar := range version.Outputs {
		if _, ok := initializedOutputs[outputVar.ID]; ok {
			continue
		}

		svo, err := TerraformClient().StateVersionOutputs.Read(ctx, outputVar.ID)
		if err != nil {
			return fmt.Errorf("error getting state version output \"%s\": %w", outputVar.ID, err)
		}

		switch svo.Name {
		case terraformOutputs.BackendAddr:
			backendAddr = svo.Value
			initializedOutputs[outputVar.ID] = struct{}{}
		default:
			// Do nothing
		}

		if svo.Name == stopAfter {
			return nil
		}
	}

	return nil
}

func InitBackendAddr(ctx context.Context) error {
	mg.Deps(InitWorkspace)

	if backendAddr != "" {
		return nil
	}

	if Workspace() == Local {
		backendAddr = "http://localhost:8080"
		return nil
	}

	initOutputVars(ctx, terraformOutputs.BackendAddr)

	if backendAddr != "" {
		return fmt.Errorf("could not initialize backendAddr")
	}

	return nil
}

func BackendAddr() string {
	if backendAddr == "" {
		panic("backend addr not initialized")
	}

	return backendAddr
}

func TerraformVars() (*TerraformVarType, error) {
	mg.Deps(InitWorkspace)

	if Workspace() == Local {
		return nil, errors.New("no local terraform vars, use \"testing\"")
	}

	return TerraformVarMap[Workspace()], nil
}
