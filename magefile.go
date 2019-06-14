// +build mage

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/shurcooL/vfsgen"
	// mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Deploy() error {
	mg.Deps(TemplatesBindata)
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
	cmd := exec.Command("hugo")
	cmd.Dir = "static"
	if mg.Verbose() {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DeployStatic() error {
	mg.Deps(BuildStatic)
	fmt.Println("Deploying static site")
	return sh.Run("gsutil", "-m", "rsync", "-R", "static/public", "gs://daytonswingsmackdown.com")
}
