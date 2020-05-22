// +build mage

package main

import (
	"fmt"

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
	return sh.Run("toolbox", "do", "--", "protoc", "--twirp_out=.", "--go_out=.", "dynamic/rpc/dss/registration.proto")
}
