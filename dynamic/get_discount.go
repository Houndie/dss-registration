package dynamic

import (
	"encoding/json"
	"net/http"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/registration/getdiscount"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type getDiscountResponseDiscount struct {
	AppliedTo string `json:"applied_to"`
	Type      string `json:"type"`
	Percent   string `json:"percent,omitempty"`
	Dollar    int    `json:"dollar,omitempty"`
}

type getDiscountResponse struct {
	Discount []*getDiscountResponseDiscount `json:"discount,omitempty"`
	Errors   []*jsonError                   `json:"errors,omitempty"`
}

func writeGetDiscountResp(w http.ResponseWriter, logger *logrus.Logger, discount []*getDiscountResponseDiscount, errors []*jsonError) {
	resp := &getDiscountResponse{
		Discount: discount,
		Errors:   errors,
	}
	bytes, err := json.Marshal(&resp)
	if err != nil {
		logger.WithError(err).Error("Error marshalling get discount response")
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		logger.WithError(err).Error("Error writing get discount response")
		return
	}
}

func GetDiscount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Get Discount (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Get Discount")

	values := struct {
		Code string `schema:"code"`
	}{}

	err := decoder.Decode(&values, r.URL.Query())
	if err != nil {
		logger.WithError(err).Debug("error decoding url values")
		switch e := err.(type) {
		case schema.EmptyFieldError:
			writeGetDiscountResp(w, logger, nil, []*jsonError{missingParameterError(e.Key)})
			return
		case schema.MultiError:
			if len(e) == 0 {
				writeGetDiscountResp(w, logger, nil, []*jsonError{internalServerError()})
				return
			}

			var ise *jsonError
			errors := []*jsonError{}
			for _, ee := range e {
				switch eee := ee.(type) {
				case schema.EmptyFieldError:
					errors = append(errors, missingParameterError(eee.Key))
				default:
					ise = internalServerError()
				}
			}
			if ise != nil {
				errors = append(errors, ise)
			}

			writeGetDiscountResp(w, logger, nil, errors)
			return
		default:
			writeGetDiscountResp(w, logger, nil, []*jsonError{internalServerError()})
			return
		}
	}

	discounts, err := getDiscountService.GetDiscount(r.Context(), values.Code)
	if err != nil {
		switch errors.Cause(err).(type) {
		case getdiscount.ErrDiscountDoesNotExist:
			writeGetDiscountResp(w, logger, nil, []*jsonError{badParameterError("code", values.Code, "discount with this name does not exist")})
		default:
			writeGetDiscountResp(w, logger, nil, []*jsonError{internalServerError()})
		}
		return
	}

	respDiscounts := make([]*getDiscountResponseDiscount, len(discounts))
	for i, discount := range discounts {
		respDiscount := &getDiscountResponseDiscount{}
		switch discount.AppliedTo {
		case common.FullWeekendPurchaseItem:
			respDiscount.AppliedTo = "Full Weekend"
		case common.DanceOnlyPurchaseItem:
			respDiscount.AppliedTo = "Dance Only"
		case common.MixAndMatchPurchaseItem:
			respDiscount.AppliedTo = "Mix And Match"
		case common.SoloJazzPurchaseItem:
			respDiscount.AppliedTo = "Solo Jazz"
		case common.TeamCompetitionPurchaseItem:
			respDiscount.AppliedTo = "Team Competition"
		case common.TShirtPurchaseItem:
			respDiscount.AppliedTo = "TShirt"
		default:
			logger.Errorf("Unknown discount applied to %v", discount.AppliedTo)
			writeGetDiscountResp(w, logger, nil, []*jsonError{internalServerError()})
		}

		switch t := discount.ItemDiscount.(type) {
		case *common.PercentDiscount:
			respDiscount.Type = "percent"
			respDiscount.Percent = t.Amount
		case *common.DollarDiscount:
			respDiscount.Type = "dollar"
			respDiscount.Dollar = t.Amount
		default:
			logger.Error("Unknown discount type")
			writeGetDiscountResp(w, logger, nil, []*jsonError{internalServerError()})
		}

		respDiscounts[i] = respDiscount
	}
	writeGetDiscountResp(w, logger, respDiscounts, nil)
}
