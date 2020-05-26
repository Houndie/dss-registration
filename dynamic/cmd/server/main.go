package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Houndie/dss-registration/dynamic/api"
	api_discount "github.com/Houndie/dss-registration/dynamic/api/discount"
	api_registration "github.com/Houndie/dss-registration/dynamic/api/registration"
	"github.com/Houndie/dss-registration/dynamic/authorizer/google"
	"github.com/Houndie/dss-registration/dynamic/discount"
	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
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

type logHttp struct {
	logger *logrus.Logger
}

func (l *logHttp) RoundTrip(r *http.Request) (*http.Response, error) {
	l.logger.Info(r.URL)
	l.logger.Info(r.Header)
	return http.DefaultTransport.RoundTrip(r)
}

func run() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error parsing viper config: %w", err)
	}
	logger := logrus.New()

	var squareEnvironment square.Environment
	switch viper.GetString("environment") {
	case "production":
		squareEnvironment = square.Production
	case "development":
		squareEnvironment = square.Sandbox
	default:
		return fmt.Errorf("unknown environment: %s", viper.GetString("environment"))
	}
	squareClient, err := square.NewClient(viper.GetString("square_key"), squareEnvironment, &http.Client{
		Transport: &logHttp{
			logger: logger,
		},
	})
	if err != nil {
		return fmt.Errorf("error creating square client: %w", err)
	}
	authorizer := google.NewAuthorizer(&http.Client{})
	pool, err := pgxpool.Connect(context.Background(), viper.GetString("storage_dsn"))
	if err != nil {
		return fmt.Errorf("error making postgres connection")
	}
	store := postgres.NewStore(pool)
	mailClient := sendgrid.NewSendClient(viper.GetString("mail_key"))

	mux := http.NewServeMux()

	registrationService := registration.NewService(true, logger, squareClient, authorizer, store, mailClient)
	registrationServer := api_registration.NewServer(registrationService)
	registrationHandler := pb.NewRegistrationServer(registrationServer, nil)
	mux.Handle(pb.RegistrationPathPrefix, registrationHandler)

	discountService := discount.NewService(store, squareClient, logger, authorizer)
	discountServer := api_discount.NewServer(discountService)
	discountHandler := pb.NewDiscountServer(discountServer, nil)
	mux.Handle(pb.DiscountPathPrefix, discountHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type", "Twirp-Version"},
	})
	corsHandler.Log = logger

	http.ListenAndServe(":80",
		api.WithLogRequest(logger,
			corsHandler.Handler(
				api.WithAuthHandler(mux))))
	return nil
}
