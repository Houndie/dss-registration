package dynamic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/sirupsen/logrus"
)

type addRegistrationResp struct {
	CheckoutUrl string       `json:"checkout_url,omitempty"`
	Errors      []*jsonError `json:"errors,omitempty"`
}

func writeAddRegistrationResp(w http.ResponseWriter, logger *logrus.Logger, checkoutUrl string, errors []*jsonError) {
	resp := &addRegistrationResp{
		CheckoutUrl: checkoutUrl,
		Errors:      errors,
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

func AddRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Add Registration (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Add Registration")

	auth := r.Header.Get("Authorization")
	authToken := ""
	if auth != "" {
		if !strings.HasPrefix(auth, "Bearer ") {
			logger.Debug("malformed auth header")
			writeAddRegistrationResp(w, logger, "", []*jsonError{internalServerError()})
			return
		}
		authToken = strings.TrimPrefix(auth, "Bearer ")
	}

	inputs := struct {
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		Address         string `json:"address"`
		City            string `json:"city"`
		State           string `json:"state"`
		Zip             string `json:"zip"`
		Email           string `json:"email"`
		HomeScene       string `json:"home_scene"`
		Student         bool   `json:"student"`
		WeekendPassType string `json:"weekend_pass_type"`
		FullWeekend     *struct {
			Level string `json:"level"`
			Tier  int    `json:"tier"`
		} `json:"full_weekend"`
		MixAndMatch     bool   `json:"mix_and_match"`
		MixAndMatchRole string `json:"mix_and_match_role"`
		SoloJazz        bool   `json:"solo_jazz"`
		TeamCompetition bool   `json:"team_competition"`
		TeamName        string `json:"team_name"`
		TShirt          bool   `json:"tshirt"`
		TShirtSize      string `json:"tshirt_size"`
		HousingStatus   string `json:"housing_status"`
		ProvideHousing  *struct {
			MyPets           string `json:"my_pets"`
			HousingNumber    int    `json:"housing_number"`
			MyHousingDetails string `json:"my_housing_details"`
		} `json:"provide_housing"`
		RequireHousing *struct {
			PetAllergies          string `json:"pet_allergies"`
			HousingRequestDetails string `json:"housing_request_details"`
		} `json:"require_housing"`
		RedirectUrl string `json:"redirect_url"`
	}{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Warn("Error reading request body")
		writeAddRegistrationResp(w, logger, "", []*jsonError{internalServerError()})
		return
	}

	err = json.Unmarshal(body, &inputs)
	if err != nil {
		logger.WithError(err).Warn("Error unmarshaling registration form")
		writeAddRegistrationResp(w, logger, "", []*jsonError{internalServerError()})
		return
	}

	var passType common.PassType
	switch inputs.WeekendPassType {
	case "Full":
		if inputs.FullWeekend == nil {
			logger.Warnf("No data submitted for a full weekend pass")
			writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("full_weekend")})
			return

		}

		var level common.WeekendPassLevel
		switch inputs.FullWeekend.Level {
		case "Level 1":
			level = common.WeekendPassLevel1
		case "Level 2":
			level = common.WeekendPassLevel2
		case "Level 3":
			level = common.WeekendPassLevel3
		case "":
			logger.Warnf("No level submitted for a full weekend pass")
			writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("full_weekend.level")})
			return
		default:
			logger.Warnf("Could not parse workshop level %s", inputs.FullWeekend.Level)
			writeAddRegistrationResp(w, logger, "", []*jsonError{badParameterError("full_weekend.level", inputs.FullWeekend.Level, `must be "Level 1", "Level 2", or "Level 3"`)})
			return
		}

		if inputs.FullWeekend.Tier < 1 || inputs.FullWeekend.Tier > 5 {
			logger.Warnf("Found invalid workshop tier %v", inputs.FullWeekend.Tier)
			writeAddRegistrationResp(w, logger, "", []*jsonError{badParameterError("full_weekend.tier", string(inputs.FullWeekend.Tier), "must be between 1 and 5 (inclusive)")})
			return
		}
		passType = &common.WeekendPass{
			Level: level,
			Tier:  common.WeekendPassTier(inputs.FullWeekend.Tier),
		}
	case "Dance":
		passType = &common.DanceOnlyPass{}
	case "None":
		passType = &common.NoPass{}
	case "":
		logger.Warnf("No pass type submitted")
		writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("weekend_pass_type")})
		return
	default:
		logger.Warnf("Could not parse weekend pass type %s", inputs.WeekendPassType)
		writeAddRegistrationResp(w, logger, "", []*jsonError{badParameterError("weekend_pass_type", inputs.WeekendPassType, `must be "Full", "Dance", or "None"`)})
		return
	}

	var mixAndMatch *common.MixAndMatch
	if inputs.MixAndMatch {
		switch inputs.MixAndMatchRole {
		case "Leader":
			mixAndMatch = &common.MixAndMatch{
				Role: common.MixAndMatchRoleLeader,
			}
		case "Follower":
			mixAndMatch = &common.MixAndMatch{
				Role: common.MixAndMatchRoleFollower,
			}
		case "":
			logger.Warnf("Mix and match role not provided")
			writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("mix_and_match_role")})
			return
		default:
			logger.Warnf("Bad mix and match role %v", inputs.MixAndMatchRole)
			writeAddRegistrationResp(w, logger, "", []*jsonError{badParameterError("mix_and_match_role", inputs.MixAndMatchRole, `must be "Lead" or "Follow"`)})
			return
		}
	}

	var teamCompetition *common.TeamCompetition
	if inputs.TeamCompetition {
		if inputs.TeamName == "" {
			logger.Warnf("Team name not provided")
			writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("team_name")})
			return
		}
		teamCompetition = &common.TeamCompetition{
			Name: inputs.TeamName,
		}
	}

	var tShirt *common.TShirt
	if inputs.TShirt {
		switch common.TShirtStyle(inputs.TShirtSize) {
		case common.TShirtStyleUnisexS, common.TShirtStyleUnisexM, common.TShirtStyleUnisexL, common.TShirtStyleUnisexXL, common.TShirtStyleUnisex2XL, common.TShirtStyleUnisex3XL, common.TShirtStyleBellaS, common.TShirtStyleBellaM, common.TShirtStyleBellaL, common.TShirtStyleBellaXL, common.TShirtStyleBella2XL:
			tShirt = &common.TShirt{
				Style: common.TShirtStyle(inputs.TShirtSize),
			}
		case "":
			logger.Warn("No T-shirt size submitted?")
			writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("tshirt_size")})
			return
		default:
			logger.Warnf("Could not parse tshirt style %s", inputs.TShirtSize)
			writeAddRegistrationResp(w, logger, "", []*jsonError{badParameterError("tshirt_size", inputs.TShirtSize, `must be "Unisex S", "Unisex M", "Unisex L", "Unisex XL", "Unisex 2XL", "Unisex 3XL", "Bella S", "Bella M", "Bella L", "Bella XL", or "Bella 2XL"`)})
			return
		}
	}

	var housing common.Housing
	switch inputs.HousingStatus {
	case "None":
		housing = &common.NoHousing{}
	case "Require":
		housing = &common.RequireHousing{
			PetAllergies: inputs.RequireHousing.PetAllergies,
			Details:      inputs.RequireHousing.HousingRequestDetails,
		}
	case "Provide":
		housing = &common.ProvideHousing{
			Pets:     inputs.ProvideHousing.MyPets,
			Quantity: inputs.ProvideHousing.HousingNumber,
			Details:  inputs.ProvideHousing.MyHousingDetails,
		}
	case "":
		logger.Warn("No housing status")
		writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("housing_status")})
		return
	default:
		logger.Warnf("Could not parse housing status %s", inputs.HousingStatus)
		writeAddRegistrationResp(w, logger, "", []*jsonError{badParameterError("housing_status", inputs.HousingStatus, `must be "None", "Require", or "Provide"`)})
		return
	}

	if inputs.FirstName == "" {
		logger.Warn("No first name")
		writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("first_name")})
		return
	}

	if inputs.LastName == "" {
		logger.Warn("No last name")
		writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("last_name")})
		return
	}

	if inputs.Email == "" {
		logger.Warn("No email")
		writeAddRegistrationResp(w, logger, "", []*jsonError{missingParameterError("email")})
		return
	}

	url, err := addService.Add(r.Context(), &add.Registration{
		FirstName:       inputs.FirstName,
		LastName:        inputs.LastName,
		StreetAddress:   inputs.Address,
		City:            inputs.City,
		State:           inputs.State,
		ZipCode:         inputs.Zip,
		Email:           inputs.Email,
		HomeScene:       inputs.HomeScene,
		IsStudent:       inputs.Student,
		PassType:        passType,
		MixAndMatch:     mixAndMatch,
		SoloJazz:        inputs.SoloJazz,
		TeamCompetition: teamCompetition,
		TShirt:          tShirt,
		Housing:         housing,
	}, inputs.RedirectUrl, authToken)
	if err != nil {
		logger.WithError(err).Error("Error adding regitration to backend")
		writeAddRegistrationResp(w, logger, "", []*jsonError{internalServerError()})
		return
	}
	writeAddRegistrationResp(w, logger, url, nil)
}
