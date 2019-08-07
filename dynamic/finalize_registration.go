package dynamic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Houndie/dss-registration/dynamic/registration/finalize"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func writeFinalizeRegistrationResp(w http.ResponseWriter, logger *logrus.Logger, errors []*jsonError) {
	resp := &addRegistrationResp{
		Errors: errors,
	}
	bytes, err := json.Marshal(&resp)
	if err != nil {
		logger.WithError(err).Error("Error marshalling finalize registration response")
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		logger.WithError(err).Error("Error writing finalize registration response")
		return
	}
}

func FinalizeRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Finalize Registration (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Finalize Registration")

	inputs := struct {
		ReferenceId   string `json:"reference_id"`
		TransactionId string `json:"transaction_id"`
	}{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Warn("Error reading request body")
		writeFinalizeRegistrationResp(w, logger, []*jsonError{internalServerError()})
		return
	}

	err = json.Unmarshal(body, &inputs)
	if err != nil {
		logger.WithError(err).Warn("Error unmarshaling registration form")
		writeFinalizeRegistrationResp(w, logger, []*jsonError{internalServerError()})
		return
	}

	referenceId, err := uuid.FromString(inputs.ReferenceId)
	if err != nil {
		logger.WithError(err).Debug("could not decode reference id to uuid")
		writeFinalizeRegistrationResp(w, logger, []*jsonError{
			badParameterError("reference_id", inputs.ReferenceId, "could not decode reference id to uuid"),
		})
		return
	}

	err = finalizeService.Finalize(r.Context(), referenceId, inputs.TransactionId)
	if err != nil {
		switch errors.Cause(err).(type) {
		case finalize.ErrReferenceIdNotEqual:
			logger.WithError(err).Debug("reference id not equal error found")
			writeFinalizeRegistrationResp(w, logger, []*jsonError{
				badParameterError("reference_id", inputs.ReferenceId, err.Error()),
			})
			return
		default:
			logger.WithError(err).Error("error found in finalizing registration")
			writeFinalizeRegistrationResp(w, logger, []*jsonError{internalServerError()})
			return
		}
	}
	writeFinalizeRegistrationResp(w, logger, nil)

}
