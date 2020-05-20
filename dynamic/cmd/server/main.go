package main

import (
	"fmt"
	"net/http"
	"os"

	api_registration "github.com/Houndie/dss-registration/dynamic/api/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	server := &api_registration.Server{}
	twirpHandler := pb.NewRegistrationServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
	return nil
}
