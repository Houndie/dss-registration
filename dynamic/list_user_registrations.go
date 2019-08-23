package dynamic

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Houndie/dss-registration/dynamic/registration/listbyuser"
	"github.com/sirupsen/logrus"
)

type listUserRegistrationsData struct {
	Id        string    `json:"registration_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Paid      bool      `json:"paid"`
}

type listUserRegistrationsResponse struct {
	Registrations []*listUserRegistrationsData `json:"registrations,omitempty"`
	Errors        []*jsonError                 `json:"errors,omitempty"`
}

func writeListUserRegistrationsResponse(w http.ResponseWriter, logger *logrus.Logger, registrations []*listbyuser.Registration, errors []*jsonError) {
	logger.Tracef("writing %d registrations and %d errors to response", len(registrations), len(errors))
	var respReg []*listUserRegistrationsData
	if registrations != nil {
		respReg = make([]*listUserRegistrationsData, len(registrations))
		for i, r := range registrations {
			respReg[i] = &listUserRegistrationsData{
				Id:        r.Id,
				FirstName: r.FirstName,
				LastName:  r.LastName,
				Email:     r.Email,
				CreatedAt: r.CreatedAt,
				Paid:      r.Paid,
			}
		}
	}
	resp := &listUserRegistrationsResponse{
		Registrations: respReg,
		Errors:        errors,
	}
	bytes, err := json.Marshal(&resp)
	if err != nil {
		logger.WithError(err).Error("Error marshalling add registration response")
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		logger.WithError(err).Error("Error writing add registration response")
		return
	}
}

func ListUserRegistrations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("List User Registrations (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("List User Registrations")

	auth := r.Header.Get("Authorization")
	if auth == "" {
		logger.Debug("authorization header not provided")
		writeListUserRegistrationsResponse(w, logger, nil, []*jsonError{unauthorizedError()})
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		logger.Debug("malformed auth header")
		writeListUserRegistrationsResponse(w, logger, nil, []*jsonError{internalServerError()})
		return
	}
	authToken := strings.TrimPrefix(auth, "Bearer ")

	res, err := listByUserService.ListByUser(r.Context(), authToken)
	if err != nil {
		writeListUserRegistrationsResponse(w, logger, nil, []*jsonError{internalServerError()})
	}
	writeListUserRegistrationsResponse(w, logger, res, nil)
}
