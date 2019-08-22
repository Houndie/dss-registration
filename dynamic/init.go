package dynamic

import (
	"bytes"
	"context"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/authorizer/google"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/Houndie/dss-registration/dynamic/registration/finalize"
	"github.com/Houndie/dss-registration/dynamic/registration/populate"
	"github.com/Houndie/dss-registration/dynamic/square"
	storage "github.com/Houndie/dss-registration/dynamic/storage/datastore"
	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
)

const (
	SQUARE_API_KEY_CONFIG_KEY = "projects/dayton-smackdown-test/configs/registration/variables/square_key"
	LOG_LEVEL                 = "projects/dayton-smackdown-test/configs/registration/variables/log_level"
	LOG_TRACE                 = "TRACE"
	LOG_DEBUG                 = "DEBUG"
	LOG_INFO                  = "INFO"
	LOG_WARN                  = "WARN"
	LOG_ERROR                 = "ERROR"
	LOG_FATAL                 = "FATAL"
	LOG_PANIC                 = "PANIC"
)

var (
	logger          *logrus.Logger
	decoder         *schema.Decoder
	populateService *populate.Service
	addService      *add.Service
	finalizeService *finalize.Service
)

func init() {
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

	loglevel, err := service.Projects.Configs.Variables.Get(LOG_LEVEL).Do()
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

	squarekey, err := service.Projects.Configs.Variables.Get(SQUARE_API_KEY_CONFIG_KEY).Do()
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

	populateService = populate.NewService(logger, squareClient)
	addService = add.NewService(logger, store, squareClient, authorizer)
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
