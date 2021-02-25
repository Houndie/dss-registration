// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Tools() error {
	fmt.Println("syncing tools")
	return sh.Run("toolbox", "sync")
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
	cmd := exec.Command("toolbox", "do", "--", "protoc", "--proto_path", "rpc/dss", "--twirp_out=dynamic/", "--go_out=dynamic/", "--twirp_typescript_out=library=pbjs:static/gatsby/src/rpc", "registration.proto", "discount.proto", "forms.proto")
	cmd.Stderr = os.Stderr
	//cmd.Dir = "dynamic"
	err := cmd.Run()
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
