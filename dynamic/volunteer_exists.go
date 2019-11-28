package dynamic

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type volunteerExistsResponse struct {
	Exists bool         `json:"exists"`
	Errors []*jsonError `json:"errors,omitempty"`
}

func writeVolunteerExistsResp(w http.ResponseWriter, logger *logrus.Logger, exists bool, errors []*jsonError) {
	resp := &volunteerExistsResponse{
		Exists: exists,
		Errors: errors,
	}
	bytes, err := json.Marshal(&resp)
	if err != nil {
		logger.WithError(err).Error("Error marshalling add discount response")
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		logger.WithError(err).Error("Error writing add discount response")
		return
	}
}

func VolunteerExists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Volunteer Exists (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Volunteer Exists")

	auth := r.Header.Get("Authorization")
	if auth == "" {
		writeVolunteerExistsResp(w, logger, false, []*jsonError{missingParameterError("authorization header")})
		return
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		logger.Debug("malformed auth header")
		writeVolunteerExistsResp(w, logger, false, []*jsonError{internalServerError()})
		return
	}
	authToken := strings.TrimPrefix(auth, "Bearer ")

	exists, err := volunteerService.Exists(r.Context(), authToken)
	if err != nil {
		writeVolunteerExistsResp(w, logger, false, []*jsonError{internalServerError()})
		return
	}
	writeVolunteerExistsResp(w, logger, exists, nil)
}
