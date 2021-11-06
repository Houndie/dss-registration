package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Houndie/dss-registration/dynamic/api"
	api_discount "github.com/Houndie/dss-registration/dynamic/api/discount"
	api_forms "github.com/Houndie/dss-registration/dynamic/api/forms"
	api_info "github.com/Houndie/dss-registration/dynamic/api/info"
	api_registration "github.com/Houndie/dss-registration/dynamic/api/registration"
	"github.com/Houndie/dss-registration/dynamic/authorizer/auth0"
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/discount"
	"github.com/Houndie/dss-registration/dynamic/forms"
	"github.com/Houndie/dss-registration/dynamic/info"
	"github.com/Houndie/dss-registration/dynamic/object/aws"
	"github.com/Houndie/dss-registration/dynamic/recaptcha"
	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage/postgres"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/objects"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twitchtv/twirp"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
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

func init() {
	viper.SetEnvPrefix("DSS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetTypeByDefaultValue(true)
	//viper.SetDefault("admins", []string{})
	viper.AutomaticEnv()

	rootCmd.AddCommand(initCommand)
}

type corsLogger struct {
	logger *logrus.Logger
}

func (l *corsLogger) Printf(arg0 string, arg1 ...interface{}) {
	l.logger.Debugf(arg0, arg1...)
}

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "DSS backend server",
	Long:  "DSS backend server",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logrus.New()
		logger.SetLevel(logrus.TraceLevel)

		var squareEnvironment objects.Environment
		switch viper.GetString("environment") {
		case "production":
			squareEnvironment = objects.Production
		case "development":
			squareEnvironment = objects.Sandbox
		default:
			return fmt.Errorf("unknown environment: %s", viper.GetString("environment"))
		}

		squareClient, err := square.NewClient(viper.GetString("squarekey"), squareEnvironment, &http.Client{})
		if err != nil {
			return fmt.Errorf("error creating square client: %w", err)
		}
		authorizer, err := auth0.NewAuthorizer(viper.GetString("authendpoint"), &http.Client{}, logger)
		if err != nil {
			return fmt.Errorf("error creating auth0 authorizer: %w", err)
		}

		pool, err := pgxpool.Connect(context.Background(), viper.GetString("postgresurl"))
		if err != nil {
			return fmt.Errorf("error making postgres connection")
		}
		store := postgres.NewStore(pool)
		sendInBlueClient, err := sendinblue.NewClient(viper.GetString("mailkey"), &http.Client{})
		if err != nil {
			return fmt.Errorf("error making new send in blue client: %w", err)
		}

		recaptchaClient, err := recaptcha.NewClient(&http.Client{}, viper.GetString("recaptchakey"))
		if err != nil {
			return fmt.Errorf("error creating recaptcha client: %w", err)
		}

		errorHook := &twirp.ServerHooks{
			Error: func(ctx context.Context, e twirp.Error) context.Context {
				if e.Code() == twirp.Canceled ||
					e.Code() == twirp.InvalidArgument ||
					e.Code() == twirp.DeadlineExceeded ||
					e.Code() == twirp.NotFound ||
					e.Code() == twirp.BadRoute ||
					e.Code() == twirp.AlreadyExists ||
					e.Code() == twirp.PermissionDenied ||
					e.Code() == twirp.Unauthenticated ||
					e.Code() == twirp.ResourceExhausted ||
					e.Code() == twirp.FailedPrecondition ||
					e.Code() == twirp.Aborted ||
					e.Code() == twirp.OutOfRange {
					logger.Debug(e.Error())
					return ctx
				}
				logger.Error(e.Error())
				return ctx
			},
		}

		squareData := &common.SquareData{}
		if err := json.Unmarshal([]byte(viper.GetString("squaredata")), &squareData); err != nil {
			return fmt.Errorf("error unmarshalling square data: %w", err)
		}

		objectClient, err := aws.NewObjectClient(viper.GetString("aws.accesskey"), viper.GetString("aws.secretkey"), "us-east-2", viper.GetString("aws.vaxbucket"))
		if err != nil {
			return fmt.Errorf("error initializing object client: %w", err)
		}

		permissionConfig := &registration.PermissionConfig{
			List:   viper.GetString("permissions.list"),
			Update: viper.GetString("permissions.update"),
		}

		mux := http.NewServeMux()

		registrationService := registration.NewService(true, viper.GetString("environment") != "production", logger, squareClient, squareData, authorizer, store, sendInBlueClient, objectClient, permissionConfig)
		registrationServer := api_registration.NewServer(registrationService)
		registrationHandler := pb.NewRegistrationServer(registrationServer, errorHook)
		mux.Handle(pb.RegistrationPathPrefix, registrationHandler)

		discountService := discount.NewService(logger, squareData)
		discountServer := api_discount.NewServer(discountService)
		discountHandler := pb.NewDiscountServer(discountServer, errorHook)
		mux.Handle(pb.DiscountPathPrefix, discountHandler)

		formsService, err := forms.NewService(sendInBlueClient, recaptchaClient)
		if err != nil {
			return fmt.Errorf("error starting forms service: %w", err)
		}
		formsServer := api_forms.NewServer(formsService)
		formsHandler := pb.NewFormsServer(formsServer, errorHook)
		mux.Handle(pb.FormsPathPrefix, formsHandler)

		infoService := info.NewService(pool, viper.GetString("version"))
		infoServer := api_info.NewServer(infoService)
		infoHandler := pb.NewInfoServer(infoServer, errorHook)
		mux.Handle(pb.InfoPathPrefix, infoHandler)

		corsHandler := cors.New(cors.Options{
			AllowedOrigins: viper.GetStringSlice("frontend"),
			AllowedMethods: []string{"POST"},
			AllowedHeaders: []string{"Content-Type", "Twirp-Version", "Authorization"},
		})
		corsHandler.Log = &corsLogger{logger: logger}

		err = http.ListenAndServe(":"+viper.GetString("port"),
			api.WithLogRequest(logger,
				corsHandler.Handler(
					api.WithAuthHandler(mux))))
		if err != nil {
			return err
		}
		return nil
	},
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "first time setup",
	Long:  "first time setup",
	RunE: func(cmd *cobra.Command, args []string) error {
		pool, err := pgxpool.Connect(context.Background(), viper.GetString("postgresurl"))
		if err != nil {
			return fmt.Errorf("error making postgres connection")
		}
		store := postgres.NewStore(pool)
		admins := viper.GetStringSlice("admins")
		for _, admin := range admins {
			if err := store.AddAdmin(context.Background(), admin); err != nil {
				return fmt.Errorf("error adding admin: %w", err)
			}
		}
		return nil
	},
}
