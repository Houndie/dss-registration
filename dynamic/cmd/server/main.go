package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	api_registration "github.com/Houndie/dss-registration/dynamic/api/registration"
	"github.com/Houndie/dss-registration/dynamic/authorizer/google"
	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error parsing viper config: %w", err)
	}
	logger := logrus.New()
	squareClient := square.NewClient(viper.GetString("square_key"), &http.Client{})
	authorizer := google.NewAuthorizer(&http.Client{})
	pool, err := pgxpool.Connect(context.Background(), viper.GetString("storage_dsn"))
	if err != nil {
		return fmt.Errorf("error making postgres connection")
	}
	store := postgres.NewStore(pool)
	mailClient := sendgrid.NewSendClient(viper.GetString("mail_key"))
	service := registration.NewService(true, logger, squareClient, authorizer, store, mailClient)
	server := api_registration.NewServer(service)
	twirpHandler := pb.NewRegistrationServer(server, nil)

	http.ListenAndServe(":80", twirpHandler)
	return nil
}
