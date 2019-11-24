// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
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
		return "https://test.daytonswingsmackdown.com", nil
	}
	return "", fmt.Errorf("Unknown project name found (%s)", project)
}

func dynamicSite() (string, error) {
	project, err := gcpProject()
	if err != nil {
		return "", err
	}
	switch project {
	case PROJECT_PROD:
		return "https://us-central1-dayton-smackdown.cloudfunctions.net", nil
	case PROJECT_TEST:
		return "https://us-central1-dayton-smackdown-test.cloudfunctions.net", nil
	}
	return "", fmt.Errorf("Unknown project name found (%s)", project)
}

func configBase() (string, error) {
	project, err := gcpProject()
	if err != nil {
		return "", err
	}
	switch project {
	case PROJECT_PROD:
		return "projects/dayton-smackdown/configs/registration/variables/", nil
	case PROJECT_TEST:
		return "projects/dayton-smackdown-test/configs/registration/variables/", nil
	}
	return "", fmt.Errorf("Unknown project name found (%s)", project)
}

func oauthData() (string, string, error) {
	project, err := gcpProject()
	if err != nil {
		return "", "", err
	}
	switch project {
	case PROJECT_PROD:
		return "630627529793-opi2g1aqshr0e9gbujpf9s3qn97mmnhd.apps.googleusercontent.com", "AIzaSyBVc8qysybfnIkgT30KqgjulYcyIZ5ep4M", nil
	case PROJECT_TEST:
		return "166144116294-c115t8bqllktva4qp6tvjjeqe7mdggu3.apps.googleusercontent.com", "AIzaSyAJaUR7I6ADbch4OX-WdkjlYsnOrhBx3xU", nil
	}
	return "", "", fmt.Errorf("Unknown project name found (%s)", project)
}

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func DeployDynamic() error {
	mg.Deps(productionCheck)
	fmt.Println("Deploying...")

	config, err := configBase()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	functions := []string{"PopulateForm", "AddRegistration", "ListUserRegistrations", "GetUserRegistration", "UpdateRegistration", "AddDiscount", "GetDiscount"}
	errChan := make(chan error, len(functions))
	for _, function := range functions {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			errChan <- sh.Run("gcloud", "functions", "deploy", f, "--source", "dynamic", "--runtime", "go111", "--trigger-http", "--set-env-vars", fmt.Sprintf("CONFIG_ROOT=%s", config))
		}(function)
	}
	wg.Wait()
	for i := 0; i < len(functions); i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}
	return nil
}

func BuildStatic() error {
	fmt.Println("Building static site")
	sitename, err := siteName()
	if err != nil {
		return err
	}
	dynamicsite, err := dynamicSite()
	if err != nil {
		return err
	}
	clientId, apiKey, err := oauthData()
	if err != nil {
		return err
	}
	cmd := exec.Command("hugo")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "HUGO_BASEURL="+sitename)
	cmd.Env = append(cmd.Env, "HUGO_DYNAMIC="+dynamicsite)
	cmd.Env = append(cmd.Env, "HUGO_CLIENT_ID="+clientId)
	cmd.Env = append(cmd.Env, "HUGO_API_KEY="+apiKey)
	cmd.Dir = "static"
	if mg.Verbose() {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DeployStatic() error {
	mg.Deps(productionCheck, BuildStatic)
	fmt.Println("Deploying static site")
	project, err := gcpProject()
	if err != nil {
		return err
	}
	var bucketname string
	var cachecontrol string
	switch project {
	case PROJECT_PROD:
		bucketname = "gs://daytonswingsmackdown.com"
		//cachecontrol = "Cache-Control:public,max-age=3600"
		cachecontrol = "Cache-Control:private"
	case PROJECT_TEST:
		bucketname = "gs://test.daytonswingsmackdown.com"
		cachecontrol = "Cache-Control:private"
	default:
		return fmt.Errorf("Unknown project name found (%s)", project)
	}
	return sh.Run("gsutil", "-h", cachecontrol, "-m", "rsync", "-d", "-c", "-R", "static/public", bucketname)
}

func Deploy() {
	mg.Deps(DeployStatic, DeployDynamic)
	fmt.Println("All Sites Deployed")
}
