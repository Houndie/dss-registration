package dynamic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Houndie/dss-registration/dynamic/registration/adddiscount"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type addDiscountResponse struct {
	Errors []*jsonError `json:"errors,omitempty"`
}

func writeAddDiscountResp(w http.ResponseWriter, logger *logrus.Logger, errors []*jsonError) {
	resp := &addDiscountResponse{
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

func AddDiscount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Add Discount (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Add Discount")

	auth := r.Header.Get("Authorization")
	if auth == "" {
		writeAddDiscountResp(w, logger, []*jsonError{missingParameterError("authorization header")})
		return
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		logger.Debug("malformed auth header")
		writeAddDiscountResp(w, logger, []*jsonError{internalServerError()})
		return
	}
	authToken := strings.TrimPrefix(auth, "Bearer ")

	inputs := struct {
		Code      string `json:"code"`
		Discounts []struct {
			Name      string `json:"name"`
			AppliedTo string `json:"applied_to"`
		} `json:"discounts"`
	}{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Warn("Error reading request body")
		writeAddDiscountResp(w, logger, []*jsonError{internalServerError()})
		return
	}

	err = json.Unmarshal(body, &inputs)
	if err != nil {
		logger.WithError(err).Warn("Error unmarshaling registration form")
		writeAddDiscountResp(w, logger, []*jsonError{internalServerError()})
		return
	}

	discount := &adddiscount.Discount{
		Code:      inputs.Code,
		Discounts: make([]*adddiscount.SingleDiscount, len(inputs.Discounts)),
	}
	for i, inputDiscount := range inputs.Discounts {
		var appliedTo common.DiscountTarget
		switch inputDiscount.AppliedTo {
		case "Full Weekend":
			appliedTo = common.FullWeekendDiscountTarget
		case "Dance Only":
			appliedTo = common.DanceOnlyDiscountTarget
		case "Mix And Match":
			appliedTo = common.MixAndMatchDiscountTarget
		case "Solo Jazz":
			appliedTo = common.SoloJazzDiscountTarget
		case "Team Competition":
			appliedTo = common.TeamCompetitionDiscountTarget
		case "TShirt":
			appliedTo = common.TShirtDiscountTarget
		default:
			logger.Debugf("Found unknown applied to: %s", inputDiscount.AppliedTo)
			writeAddDiscountResp(w, logger, []*jsonError{badParameterError("discounts.applied_to", inputDiscount.AppliedTo, "must be one of the following: Full Weekend, Dance Only, Mix And Match, Solo Jazz, Team Competition, TShirt")})
		}
		discount.Discounts[i] = &adddiscount.SingleDiscount{
			Name:      inputDiscount.Name,
			AppliedTo: appliedTo,
		}
	}

	err = addDiscountService.AddDiscount(r.Context(), authToken, discount)
	if err != nil {
		switch errors.Cause(err).(type) {
		case adddiscount.ErrUnauthorized:
			logger.WithError(err).Debug("found unauthorized error")
			writeAddDiscountResp(w, logger, []*jsonError{unauthorizedError()})
		default:
			writeAddDiscountResp(w, logger, []*jsonError{internalServerError()})
		}
	}
	writeAddDiscountResp(w, logger, nil)
}
