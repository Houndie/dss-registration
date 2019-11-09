package dynamic

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/registration/getbyid"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type getUserRegistrationFullWeekendData struct {
	Level string `json:"level"`
	Tier  int    `json:"tier"`
}
type getUserRegistrationProvideHousingData struct {
	MyPets           string `json:"my_pets"`
	HousingNumber    int    `json:"housing_number"`
	MyHousingDetails string `json:"my_housing_details"`
}

type getUserRegistrationRequireHousingData struct {
	PetAllergies          string `json:"pet_allergies"`
	HousingRequestDetails string `json:"housing_request_details"`
}

type getUserRegistrationUnpaidItems struct {
	Ids   []string `json:"ids"`
	Items []string `json:"items,omitempty"`
	Cost  int      `json:"cost"`
}

type getUserRegistrationSingleDiscount struct {
	AppliedTo string `json:"applied_to"`
	Type      string `json:"type"`
	Percent   string `json:"percent,omitempty"`
	Dollar    int    `json:"dollar,omitempty"`
}

type getUserRegistrationDiscount struct {
	Code      string                               `json:"code"`
	Discounts []*getUserRegistrationSingleDiscount `json:"discounts"`
}

type getUserRegistrationsData struct {
	Id              string                                 `json:"registration_id"`
	FirstName       string                                 `json:"first_name"`
	LastName        string                                 `json:"last_name"`
	Address         string                                 `json:"address"`
	City            string                                 `json:"city"`
	State           string                                 `json:"state"`
	Zip             string                                 `json:"zip"`
	Email           string                                 `json:"email"`
	HomeScene       string                                 `json:"home_scene"`
	Student         bool                                   `json:"student"`
	WeekendPassType string                                 `json:"weekend_pass_type"`
	FullWeekend     *getUserRegistrationFullWeekendData    `json:"full_weekend,omitempty"`
	MixAndMatch     bool                                   `json:"mix_and_match"`
	MixAndMatchRole string                                 `json:"mix_and_match_role"`
	SoloJazz        bool                                   `json:"solo_jazz"`
	TeamCompetition bool                                   `json:"team_competition"`
	TeamName        string                                 `json:"team_name"`
	TShirt          bool                                   `json:"tshirt"`
	TShirtSize      string                                 `json:"tshirt_size"`
	HousingStatus   string                                 `json:"housing_status"`
	ProvideHousing  *getUserRegistrationProvideHousingData `json:"provide_housing,omitempty"`
	RequireHousing  *getUserRegistrationRequireHousingData `json:"require_housing,omitempty"`
	CreatedAt       time.Time                              `json:"created_at"`
	UnpaidItems     *getUserRegistrationUnpaidItems        `json:"unpaid_items,omitempty"`
	Discounts       []*getUserRegistrationDiscount         `json:"discounts,omitempty"`
}

type getUserRegistrationResponse struct {
	Registration *getUserRegistrationsData `json:"registration,omitempty"`
	Errors       []*jsonError              `json:"errors,omitempty"`
}

func writeGetUserRegistrationResp(w http.ResponseWriter, logger *logrus.Logger, r *getUserRegistrationsData, errors []*jsonError) {
	resp := &getUserRegistrationResponse{
		Registration: r,
		Errors:       errors,
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

func GetUserRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Get User Registration (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Get User Registration")

	auth := r.Header.Get("Authorization")
	if auth == "" {
		writeGetUserRegistrationResp(w, logger, nil, []*jsonError{missingParameterError("authorization header")})
		return
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		logger.Debug("malformed auth header")
		writeGetUserRegistrationResp(w, logger, nil, []*jsonError{internalServerError()})
		return
	}
	authToken := strings.TrimPrefix(auth, "Bearer ")

	values := struct {
		Id string `schema:"id"`
	}{}

	err := decoder.Decode(&values, r.URL.Query())
	if err != nil {
		logger.WithError(err).Debug("error decoding url values")
		switch e := err.(type) {
		case schema.EmptyFieldError:
			writeGetUserRegistrationResp(w, logger, nil, []*jsonError{missingParameterError(e.Key)})
			return
		case schema.MultiError:
			if len(e) == 0 {
				writeGetUserRegistrationResp(w, logger, nil, []*jsonError{internalServerError()})
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

			writeGetUserRegistrationResp(w, logger, nil, errors)
			return
		default:
			writeGetUserRegistrationResp(w, logger, nil, []*jsonError{internalServerError()})
			return
		}
	}

	registration, err := getByIdService.GetById(r.Context(), authToken, values.Id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case getbyid.ErrBadRegistrationId:
			writeGetUserRegistrationResp(w, logger, nil, []*jsonError{badParameterError("id", values.Id, err.Error())})
			return
		default:
			writeGetUserRegistrationResp(w, logger, nil, []*jsonError{internalServerError()})
			return
		}
	}

	resp := &getUserRegistrationsData{
		Id:        values.Id,
		FirstName: registration.FirstName,
		LastName:  registration.LastName,
		Address:   registration.StreetAddress,
		City:      registration.City,
		State:     registration.State,
		Zip:       registration.ZipCode,
		Email:     registration.Email,
		HomeScene: registration.HomeScene,
		Student:   registration.IsStudent,
		SoloJazz:  registration.SoloJazz,
		CreatedAt: registration.CreatedAt,
	}

	switch t := registration.PassType.(type) {
	case *common.WeekendPass:
		resp.WeekendPassType = "Full"

		var level string
		switch t.Level {
		case common.WeekendPassLevel1:
			level = "Level 1"
		case common.WeekendPassLevel2:
			level = "Level 2"
		case common.WeekendPassLevel3:
			level = "Level 3"
		default:
			logger.Errorf("unknown level %v", t)
			writeGetUserRegistrationResp(w, logger, nil, []*jsonError{internalServerError()})
			return
		}
		resp.FullWeekend = &getUserRegistrationFullWeekendData{
			Level: level,
			Tier:  int(t.Tier),
		}
	case *common.DanceOnlyPass:
		resp.WeekendPassType = "Dance"
	case *common.NoPass:
		resp.WeekendPassType = "None"
	default:
		logger.Error("unknown full weekend pass type found")
		writeGetUserRegistrationResp(w, logger, nil, []*jsonError{internalServerError()})
		return
	}

	if registration.MixAndMatch != nil {
		resp.MixAndMatch = true
		resp.MixAndMatchRole = string(registration.MixAndMatch.Role)
	}

	if registration.TeamCompetition != nil {
		resp.TeamCompetition = true
		resp.TeamName = registration.TeamCompetition.Name
	}

	if registration.TShirt != nil {
		resp.TShirt = true
		resp.TShirtSize = string(registration.TShirt.Style)
	}

	switch t := registration.Housing.(type) {
	case *common.ProvideHousing:
		resp.HousingStatus = "Provide"
		resp.ProvideHousing = &getUserRegistrationProvideHousingData{
			MyPets:           t.Pets,
			HousingNumber:    t.Quantity,
			MyHousingDetails: t.Details,
		}
	case *common.RequireHousing:
		resp.HousingStatus = "Require"
		resp.RequireHousing = &getUserRegistrationRequireHousingData{
			PetAllergies:          t.PetAllergies,
			HousingRequestDetails: t.Details,
		}
	case *common.NoHousing:
		resp.HousingStatus = "None"
	default:
		logger.Error("unknown housing request type found")
		writeGetUserRegistrationResp(w, logger, nil, []*jsonError{internalServerError()})
		return
	}

	if registration.UnpaidItems != nil {
		resp.UnpaidItems = &getUserRegistrationUnpaidItems{
			Ids:   registration.UnpaidItems.OrderIds,
			Items: registration.UnpaidItems.Items,
			Cost:  registration.UnpaidItems.Cost,
		}
	}

	if len(registration.Discounts) > 0 {
		discounts := make([]*getUserRegistrationDiscount, len(registration.Discounts))
		for i, discount := range registration.Discounts {
			singleDiscounts := make([]*getUserRegistrationSingleDiscount, len(discount.Discounts))
			for j, sd := range discount.Discounts {
				respDiscount := &getUserRegistrationSingleDiscount{}
				switch sd.AppliedTo {
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
					logger.Errorf("Unknown discount applied to %v", sd.AppliedTo)
					writeGetDiscountResp(w, logger, nil, []*jsonError{internalServerError()})
				}

				switch t := sd.ItemDiscount.(type) {
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

				singleDiscounts[j] = respDiscount
			}
			discounts[i] = &getUserRegistrationDiscount{
				Code:      discount.Code,
				Discounts: singleDiscounts,
			}
		}
		resp.Discounts = discounts
	}

	writeGetUserRegistrationResp(w, logger, resp, nil)
}
