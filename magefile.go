// +build mage

package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Houndie/dss-registration/mage"
	"github.com/Houndie/toolbox/pkg/toolbox"
	"github.com/magefile/mage/mg"
)

func Tools() error {
	fmt.Println("syncing tools")
	return toolbox.Sync()
}

const eslint = "eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars, strict, no-lone-blocks, default-case"

func GenerateProtoc() error {
	mg.Deps(Tools)
	fmt.Println("generating protocs")
	for _, file := range []string{"registration", "discount", "forms"} {
		pbjs := exec.Command("npx", "pbjs", "-t", "static-module", "-w", "commonjs", "-l", eslint, "-r", file, "-o", "static/gatsby/src/rpc/"+file+".pb.js", "rpc/dss/"+file+".proto")
		pbjs.Stderr = os.Stderr
		err := pbjs.Run()
		if err != nil {
			return err
		}
		pbts := exec.Command("npx", "pbts", "--no-comments", "-o", "static/gatsby/src/rpc/"+file+".pb.d.ts", "static/gatsby/src/rpc/"+file+".pb.js")
		pbts.Stderr = os.Stderr
		err = pbts.Run()
		if err != nil {
			return err
		}
	}
	cmd, err := toolbox.Command("protoc", "--proto_path", "rpc/dss", "--twirp_out=dynamic/", "--go_out=dynamic/", "--twirp_typescript_out=library=pbjs:static/gatsby/src/rpc", "registration.proto", "discount.proto", "forms.proto")
	if err != nil {
		return err
	}

	cmd.Stderr = os.Stderr
	//cmd.Dir = "dynamic"
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Prepend /* eslint-disable */ to files
	files, err := filepath.Glob("static/gatsby/src/rpc/*_pb.js")
	if err != nil {
		return fmt.Errorf("error globbing pb files: %w", err)
	}
	for _, file := range files {
		f, err := os.OpenFile(file, os.O_RDWR, 0755)
		if err != nil {
			return fmt.Errorf("error opening file %v: %w", file, err)
		}
		defer f.Close()

		fileText, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("error reading file %v: %w", file, err)
		}

		if _, err := f.Seek(0, 0); err != nil {
			return fmt.Errorf("error resetting seek back to start of file %v: %w", file, err)
		}

		if _, err = f.WriteString("/* eslint-disable */\n"); err != nil {
			return fmt.Errorf("error prepending disable string file %v: %w", file, err)
		}
		if _, err = f.Write(fileText); err != nil {
			return fmt.Errorf("error rewriting file %v: %w", file, err)
		}
		if err := f.Close(); err != nil {
			return fmt.Errorf("error closing file %v: %w", file, err)
		}
	}

	return nil
}

func SetTerraformDeployVersion(ctx context.Context) error {
	fmt.Println("setting terraform deploy version")
	workspaceName, ok := os.LookupEnv("TERRAFORM_WORKSPACE")
	if !ok {
		return errors.New("environment variable TERRAFORM_WORKSPACE must not be empty")
	}

	apiKey, ok := os.LookupEnv("TERRAFORM_API_KEY")
	if !ok {
		return errors.New("environment variable TERRAFORM_API_KEY must not be empty")
	}

	deployVersion, ok := os.LookupEnv("DEPLOY_VERSION")
	if !ok {
		return errors.New("environment variable DEPLOY_VERSION not set")
	}

	client := &http.Client{
		Transport: &mage.TerraformTransport{
			ApiKey: apiKey,
		},
		Timeout: 10 * time.Second,
	}

	workspaceID, err := mage.TerraformWorkspaceID(ctx, client, workspaceName)
	if err != nil {
		return err
	}

	varID, err := mage.TerraformDeployVersionID(ctx, client, workspaceID)
	if err != nil {
		return err
	}

	if err := mage.SetTerraformDeployVersionID(ctx, client, workspaceID, varID, deployVersion); err != nil {
		return err
	}

	return nil
}
