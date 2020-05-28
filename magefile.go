// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Tools() error {
	fmt.Println("syncing tools")
	return sh.Run("toolbox", "sync")
}

func GenerateProtoc() error {
	mg.Deps(Tools)
	fmt.Println("generating protocs")
	cmd := exec.Command("toolbox", "do", "--", "protoc", "--proto_path", "rpc/dss", "--twirp_out=dynamic/", "--go_out=dynamic/", "--twirp_js_out=static/rpc", "--js_out=import_style=commonjs,binary:static/rpc", "registration.proto", "discount.proto")
	cmd.Stderr = os.Stderr
	//cmd.Dir = "dynamic"
	return cmd.Run()
}

func CompileReact() error {
	fmt.Println("Compiling react components")

	cmd := exec.Command("npx", "webpack", "--mode", "development")
	cmd.Dir = "static"
	if mg.Verbose() {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error generating react component: %w", err)
	}
	return nil
}

func BuildStatic() error {
	mg.Deps(CompileReact)
	fmt.Println("Building static site")
	sitename := "http://localhost:8081"
	dynamicsite := "https://us-central1-dayton-smackdown-test.cloudfunctions.net"
	clientId, apiKey := "166144116294-c115t8bqllktva4qp6tvjjeqe7mdggu3.apps.googleusercontent.com", "AIzaSyAJaUR7I6ADbch4OX-WdkjlYsnOrhBx3xU"
	cmd := exec.Command("toolbox", "do", "--", "hugo")
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
