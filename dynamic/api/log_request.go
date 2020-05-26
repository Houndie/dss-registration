package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func WithLogRequest(logger *logrus.Logger, base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.URL.String())
		base.ServeHTTP(w, r)
	})
}
