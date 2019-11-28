package dynamic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/volunteer"
	"github.com/sirupsen/logrus"
)

type insertVolunteerResponse struct {
	Errors []*jsonError `json:"errors,omitempty"`
}

func writeInsertVolunteerResp(w http.ResponseWriter, logger *logrus.Logger, errors []*jsonError) {
	resp := &insertVolunteerResponse{
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

func InsertVolunteer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Insert Volunteer (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Insert Volunteer")

	auth := r.Header.Get("Authorization")
	if auth == "" {
		writeInsertVolunteerResp(w, logger, []*jsonError{missingParameterError("authorization header")})
		return
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		logger.Debug("malformed auth header")
		writeInsertVolunteerResp(w, logger, []*jsonError{internalServerError()})
		return
	}
	authToken := strings.TrimPrefix(auth, "Bearer ")

	inputs := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Warn("Error reading request body")
		writeInsertVolunteerResp(w, logger, []*jsonError{internalServerError()})
		return
	}

	err = json.Unmarshal(body, &inputs)
	if err != nil {
		logger.WithError(err).Warn("Error unmarshaling registration form")
		writeInsertVolunteerResp(w, logger, []*jsonError{internalServerError()})
		return
	}

	v := &volunteer.VolunteerSubmission{
		Name:  inputs.Name,
		Email: inputs.Email,
	}

	err = volunteerService.Insert(r.Context(), authToken, v)
	if err != nil {
		switch err.(type) {
		case storage.ErrVolunteerExists:
			writeInsertVolunteerResp(w, logger, []*jsonError{alreadyExistsError()})
		default:
			writeInsertVolunteerResp(w, logger, []*jsonError{internalServerError()})
		}
		return
	}
	writeInsertVolunteerResp(w, logger, nil)
}
