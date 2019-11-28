package dynamic

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/authorizer/google"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/Houndie/dss-registration/dynamic/registration/adddiscount"
	"github.com/Houndie/dss-registration/dynamic/registration/getbyid"
	"github.com/Houndie/dss-registration/dynamic/registration/getdiscount"
	"github.com/Houndie/dss-registration/dynamic/registration/listbyuser"
	"github.com/Houndie/dss-registration/dynamic/registration/populate"
	"github.com/Houndie/dss-registration/dynamic/registration/update"
	"github.com/Houndie/dss-registration/dynamic/square"
	storage "github.com/Houndie/dss-registration/dynamic/storage/datastore"
	"github.com/Houndie/dss-registration/dynamic/volunteer"
	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/gorilla/schema"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sirupsen/logrus"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
)

const (
	SQUARE_API_KEY_CONFIG_KEY = "square_key"
	MAIL_API_KEY_CONFIG_KEY   = "mail_key"
	LOG_LEVEL                 = "log_level"
	ACTIVE_CONFIG_KEY         = "active"
	LOG_TRACE                 = "TRACE"
	LOG_DEBUG                 = "DEBUG"
	LOG_INFO                  = "INFO"
	LOG_WARN                  = "WARN"
	LOG_ERROR                 = "ERROR"
	LOG_FATAL                 = "FATAL"
	LOG_PANIC                 = "PANIC"
)

var (
	logger             *logrus.Logger
	decoder            *schema.Decoder
	populateService    *populate.Service
	addService         *add.Service
	listByUserService  *listbyuser.Service
	getByIdService     *getbyid.Service
	updateService      *update.Service
	addDiscountService *adddiscount.Service
	getDiscountService *getdiscount.Service
	volunteerService   *volunteer.Service
)

func init() {
	configRoot := os.Getenv("CONFIG_ROOT")
	if configRoot == "" {
		logger.Fatalf("Config root environment variable not set")
		os.Exit(1)
	}
	logger = logrus.New()
	logger.SetFormatter(stackdriver.NewFormatter(
		stackdriver.WithService("Registration"),
		stackdriver.WithVersion("v0.0.1"),
	))

	ctx := context.Background()
	service, err := runtimeconfig.NewService(ctx)
	if err != nil {
		logger.WithError(err).Fatal("Could not access GCP runtime config")
		os.Exit(1)
	}

	loglevel, err := service.Projects.Configs.Variables.Get(configRoot + LOG_LEVEL).Do()
	if err != nil {
		logger.WithError(err).Fatal("Could not fetch log level config")
		os.Exit(1)
	}
	if loglevel.Value != "" {
		logger.WithError(err).Fatal("Log level config set as value")
		os.Exit(1)
	}
	if loglevel.Text == "" {
		logger.WithError(err).Fatal("Log level config text not found")
		os.Exit(1)
	}
	loglevelmap := map[string]logrus.Level{
		LOG_TRACE: logrus.TraceLevel,
		LOG_DEBUG: logrus.DebugLevel,
		LOG_INFO:  logrus.InfoLevel,
		LOG_WARN:  logrus.WarnLevel,
		LOG_ERROR: logrus.ErrorLevel,
		LOG_FATAL: logrus.FatalLevel,
		LOG_PANIC: logrus.PanicLevel,
	}
	level, ok := loglevelmap[loglevel.Text]
	if !ok {
		logger.Fatalf("Could not find log level %s", loglevel.Text)
		os.Exit(1)
	}
	logger.SetLevel(level)

	squarekey, err := service.Projects.Configs.Variables.Get(configRoot + SQUARE_API_KEY_CONFIG_KEY).Do()
	if err != nil {
		logger.WithError(err).Fatal("Could not fetch square api key")
		os.Exit(1)
	}
	if squarekey.Value != "" {
		logger.WithError(err).Fatal("Square API key set as value")
		os.Exit(1)
	}
	if squarekey.Text == "" {
		logger.WithError(err).Fatal("Square API key text not found")
		os.Exit(1)
	}

	mailkey, err := service.Projects.Configs.Variables.Get(configRoot + MAIL_API_KEY_CONFIG_KEY).Do()
	if err != nil {
		logger.WithError(err).Fatal("Could not fetch sendgrid api key")
		os.Exit(1)
	}
	if mailkey.Value != "" {
		logger.WithError(err).Fatal("Sendgrid API key set as value")
		os.Exit(1)
	}
	if mailkey.Text == "" {
		logger.WithError(err).Fatal("Sendgrid API key text not found")
		os.Exit(1)
	}

	activeString, err := service.Projects.Configs.Variables.Get(configRoot + ACTIVE_CONFIG_KEY).Do()
	if err != nil {
		logger.WithError(err).Fatal("could not fetch if registration is active")
		os.Exit(1)
	}
	if activeString.Value != "" {
		logger.WithError(err).Fatal("active config key set as value")
		os.Exit(1)
	}
	active, err := strconv.ParseBool(activeString.Text)
	if err != nil {
		logger.WithError(err).WithField("active", active).Fatal("could not convert active config key to bool")
		os.Exit(1)
	}

	httpClient := &http.Client{
		Transport: &logRequests{
			logger: logger,
			wrap:   http.DefaultTransport,
		},
	}
	squareClient := square.NewClient(squarekey.Text, httpClient)

	datastore, err := datastore.NewClient(ctx, datastore.DetectProjectID)
	if err != nil {
		logger.WithError(err).Fatal("Could not get datastore connection")
		os.Exit(1)
	}

	store := storage.NewDatastore(datastore)

	authorizer := google.NewAuthorizer(httpClient)

	mail := sendgrid.NewSendClient(mailkey.Text)

	populateService = populate.NewService(logger, squareClient)
	addService = add.NewService(logger, store, squareClient, authorizer, mail, active)
	listByUserService = listbyuser.NewService(authorizer, logger, store, squareClient)
	getByIdService = getbyid.NewService(logger, authorizer, store, squareClient)
	updateService = update.NewService(logger, authorizer, store, squareClient)
	addDiscountService = adddiscount.NewService(logger, store, authorizer)
	getDiscountService = getdiscount.NewService(logger, store, squareClient)
	volunteerService = volunteer.NewService(logger, store, authorizer)
	decoder = schema.NewDecoder()
}

type logRequests struct {
	logger *logrus.Logger
	wrap   http.RoundTripper
}

func (l *logRequests) RoundTrip(r *http.Request) (*http.Response, error) {
	headerCopy := http.Header{}
	for key, value := range r.Header {
		if key == "Authorization" {
			headerCopy[key] = []string{"[redacted]"}
			continue
		}
		headerCopy[key] = value
	}
	headerbuf := &bytes.Buffer{}
	headerCopy.Write(headerbuf)
	l.logger.WithFields(logrus.Fields{
		"query":   r.URL.String(),
		"headers": headerbuf.String(),
	}).Trace("Sending http request")
	return l.wrap.RoundTrip(r)
}
