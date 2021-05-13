// +build mage

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Houndie/dss-registration/mage"
	"github.com/Houndie/toolbox/pkg/toolbox"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-tfe"
	"github.com/magefile/mage/mg"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Tools() error {
	fmt.Println("syncing tools")
	return toolbox.Sync()
}

const eslint = "eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars, strict, no-lone-blocks, default-case"

type Protoc mg.Namespace

func (Protoc) Generate() error {
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

type Terraform mg.Namespace

func (Terraform) SetDeployVersion(ctx context.Context) error {
	mg.Deps(mage.InitWorkspace, mage.InitDeployVersion, mage.InitTerraformClient)
	fmt.Println("setting terraform deploy version")

	deployVersion := mage.DeployVersion()
	_, err := mage.TerraformClient().Variables.Update(ctx, mage.TerraformWorkspace[mage.Workspace()], mage.TerraformDeployVar[mage.Workspace()], tfe.VariableUpdateOptions{
		Value: &deployVersion,
	})
	if err != nil {
		return fmt.Errorf("error updating terraform deploy version: %w", err)
	}

	return nil
}

func (Terraform) Apply(ctx context.Context) error {
	mg.Deps(mage.InitTerraformClient, mage.InitWorkspace)

	autoQueueRuns := false
	configurationVersion, err := mage.TerraformClient().ConfigurationVersions.Create(
		ctx,
		mage.TerraformWorkspace[mage.Workspace()],
		tfe.ConfigurationVersionCreateOptions{
			AutoQueueRuns: &autoQueueRuns,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating terraform configuration version: %w", err)
	}

	if err := mage.TerraformClient().ConfigurationVersions.Upload(ctx, configurationVersion.UploadURL, "terraform"); err != nil {
		return fmt.Errorf("error uploading terraform files: %w", err)
	}

	timeoutCtx1, cancel1 := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel1()
	for {
		if mg.Verbose() {
			fmt.Fprintln(os.Stderr, "polling to see if configuration upload is complete")
		}

		updatedConfigurationVersion, err := mage.TerraformClient().ConfigurationVersions.Read(timeoutCtx1, configurationVersion.ID)
		if timeoutCtx1.Err() != nil {
			return fmt.Errorf("terraform configuration version did not finish uploading before poll timeout")
		}
		if err != nil {
			return fmt.Errorf("error reading configuration version")
		}

		if updatedConfigurationVersion.Status == tfe.ConfigurationErrored || updatedConfigurationVersion.Status == tfe.ConfigurationUploaded {
			if mg.Verbose() {
				fmt.Fprintln(os.Stderr, "done polling for configuration version upload")
			}

			break
		}

		time.Sleep(5 * time.Second)
	}

	run, err := mage.TerraformClient().Runs.Create(ctx, tfe.RunCreateOptions{
		ConfigurationVersion: configurationVersion,
		Workspace: &tfe.Workspace{
			ID: mage.TerraformWorkspace[mage.Workspace()],
		},
	})
	if err != nil {
		return fmt.Errorf("error creating terraform run: %w", err)
	}

	timeoutCtx2, cancel2 := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel2()
	for {
		if mg.Verbose() {
			fmt.Fprintln(os.Stderr, "polling to see if terraform run is complete")
		}

		updatedRun, err := mage.TerraformClient().Runs.Read(timeoutCtx2, run.ID)
		if timeoutCtx2.Err() != nil {
			return fmt.Errorf("terraform run did not finish applying before poll timeout")
		}

		if err != nil {
			return fmt.Errorf("error polling for terraform run: %w", err)
		}

		if updatedRun.Status == tfe.RunPlannedAndFinished || updatedRun.Status == tfe.RunApplied {
			if mg.Verbose() {
				fmt.Fprintln(os.Stderr, "run complete!")
			}

			break
		} else if updatedRun.Status == tfe.RunPlanned {
			if mg.Verbose() {
				fmt.Fprintln(os.Stderr, "run state \"planned\" found, attempting to apply")
			}

			if err := mage.TerraformClient().Runs.Apply(ctx, run.ID, tfe.RunApplyOptions{}); err != nil {
				return fmt.Errorf("error applying terraform run: %w", err)
			}
		}

		if mg.Verbose() {
			fmt.Fprintf(os.Stderr, "found status %v\n", updatedRun.Status)
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

type Docker mg.Namespace

func (Docker) Build(ctx context.Context) error {
	mg.Deps(mage.InitDeployVersion, mage.InitWorkspace, mage.InitDockerClient)
	client := mage.DockerClient()

	if err := mage.DockerBuild(ctx, client, mage.HerokuProject[mage.Workspace()], mage.DeployVersion()); err != nil {
		return err
	}

	return nil
}

func (Docker) Save(ctx context.Context) error {
	mg.Deps(mage.InitDeployVersion, mage.InitWorkspace, mage.InitDockerClient, mage.InitDockerCache)
	client := mage.DockerClient()

	if err := mage.DockerSave(ctx, client, mage.HerokuProject[mage.Workspace()], mage.DeployVersion(), mage.DockerCache()); err != nil {
		return err
	}

	return nil
}

func (Docker) Load(ctx context.Context) error {
	mg.Deps(mage.InitDockerClient, mage.InitDockerCache)
	client := mage.DockerClient()

	if err := mage.DockerLoad(ctx, client, mage.DockerCache()); err != nil {
		return err
	}

	return nil
}

func (Docker) DeployHeroku(ctx context.Context) error {
	mg.Deps(mage.InitHerokuAPIKey, mage.InitDeployVersion, mage.InitWorkspace, mage.InitDockerClient)

	client := mage.DockerClient()

	if err := mage.DockerPush(ctx, client, mage.HerokuAPIKey(), mage.HerokuProject[mage.Workspace()], mage.DeployVersion()); err != nil {
		return err
	}

	return nil
}

func Migrate(ctx context.Context) error {
	mg.Deps(mage.InitMigrationURL)

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %w", err)
	}
	m, err := migrate.New("file://"+filepath.Join(cwd, "dynamic/storage/postgres/migrations"), mage.MigrationURL())
	if err != nil {
		return fmt.Errorf("error creating migration client: %w", err)
	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("error migrating up: %w", err)
	}

	return nil
}
