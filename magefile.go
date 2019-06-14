// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/shurcooL/vfsgen"
)

const (
	PROJECT_PROD = "dayton-smackdown"
	PROJECT_TEST = "dayton-smackdown-test"
)

func gcpProject() (string, error) {
	cmd := exec.Command("gcloud", "config", "list", "--format", "value(core.project)")
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

func productionCheck() error {
	project, err := gcpProject()
	if err != nil {
		return err
	}
	if project == PROJECT_PROD {
		fmt.Println("Attempting to do something that will affect production!")
		fmt.Println("Type \"production\" to continue")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		t := scan.Text()
		if t != "production" {
			return fmt.Errorf("Production check not passed! (found %s)", t)
		}
	}
	return nil
}

func siteName() (string, error) {
	project, err := gcpProject()
	if err != nil {
		return "", err
	}
	switch project {
	case PROJECT_PROD:
		return "https://daytonswingsmackdown.com", nil
	case PROJECT_TEST:
		return "http://test.daytonswingsmackdown.com", nil
	}
	return "", fmt.Errorf("Unknown project name found (%s)", project)
}

func bucketName() (string, error) {
	project, err := gcpProject()
	if err != nil {
		return "", err
	}
	switch project {
	case PROJECT_PROD:
		return "gs://daytonswingsmackdown.com", nil
	case PROJECT_TEST:
		return "gs://test.daytonswingsmackdown.com", nil
	}
	return "", fmt.Errorf("Unknown project name found (%s)", project)
}

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func DeployDynamic() error {
	mg.Deps(productionCheck, TemplatesBindata)
	fmt.Println("Deploying...")
	return sh.Run("gcloud", "functions", "deploy", "Registration", "--source", "src", "--runtime", "go111", "--trigger-http")
}

func TemplatesBindata() error {
	fmt.Println("Generating Bindata")
	return vfsgen.Generate(http.Dir("src/templates"), vfsgen.Options{
		Filename:     "src/templates/vfsdata.go",
		PackageName:  "templates",
		VariableName: "Assets",
	})
}

func BuildStatic() error {
	fmt.Println("Building static site")
	sitename, err := siteName()
	if err != nil {
		return err
	}
	cmd := exec.Command("hugo")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "HUGO_BASEURL="+sitename)
	cmd.Dir = "static"
	if mg.Verbose() {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DeployStatic() error {
	mg.Deps(productionCheck, BuildStatic)
	bucketname, err := bucketName()
	if err != nil {
		return err
	}
	fmt.Println("Deploying static site")
	return sh.Run("gsutil", "-m", "rsync", "-R", "static/public", bucketname)
}
