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

func GenerateServerProtoc() error {
	mg.Deps(Tools)
	fmt.Println("generating server protocs")
	cmd := exec.Command("toolbox", "do", "--", "protoc", "--twirp_out=.", "--go_out=.", "rpc/dss/registration.proto", "rpc/dss/discount.proto")
	cmd.Stderr = os.Stderr
	cmd.Dir = "dynamic"
	return cmd.Run()
}
