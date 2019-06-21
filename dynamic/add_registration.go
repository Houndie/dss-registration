package dynamic

import (
	"fmt"
	"net/http"

	"github.com/Houndie/dss-registration/dynamic/registration/add"
)

func AddRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Add Registration (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Add Registration")
	err := r.ParseForm()
	if err != nil {
		logger.WithError(err).Warn("Error parsing registration form")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	formData := struct {
		FirstName             string `schema:"first_name"`
		LastName              string `schema:"last_name"`
		Address               string `schema:"address"`
		City                  string `schema:"city"`
		State                 string `schema:"state"`
		Zip                   string `schema:"zip"`
		Email                 string `schema:"email"`
		HomeScene             string `schema:"home_scene"`
		WeekendPassType       string `schema:"weekend_pass_type"`
		Student               bool   `schema:"student"`
		WorkshopLevel         string `schema:"workshop_level"`
		MixAndMatch           bool   `schema:"mix_and_match"`
		MixAndMatchRole       string `schema:"mix_and_match_role"`
		SoloJazz              bool   `schema:"solo_jazz"`
		TeamCompetition       bool   `schema:"team_competition"`
		TeamName              string `schema:"team_name"`
		TShirt                bool   `schema:"tshirt"`
		TShirtSize            string `schema:"tshirt_size"`
		HousingStatus         string `schema:"housing_status"`
		MyPets                string `schema:"my_pets"`
		HousingNumber         int    `schema:"housing_number"`
		MyHousingDetails      string `schema:"my_housing_details"`
		PetAllergies          string `schema:"pet_allergies"`
		HousingRequestDetails string `schema:"housing_request_details"`
	}{}

	err = decoder.Decode(&formData, r.PostForm)
	if err != nil {
		fmt.Println(r.PostForm)
		logger.WithError(err).Warn("Error decoding registration form")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var passType add.PassType
	switch formData.WeekendPassType {
	case "Full":
		var level add.WeekendPassLevel
		switch formData.WorkshopLevel {
		case "Level 1":
			level = add.Level1
		case "Level 2":
			level = add.Level2
		case "Level 3":
			level = add.Level3
		case "None":
			logger.Warnf("No level submitted for a weekend pass?")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		default:
			logger.Warnf("Could not parse workshop level %s", formData.WorkshopLevel)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		passType = &add.WeekendPass{
			Level: level,
		}
	case "Dance":
		passType = &add.DanceOnlyPass{}
	case "None":
		passType = &add.NoPass{}
	default:
		logger.Warnf("Could not parse weekend pass type %s", formData.WeekendPassType)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var mixAndMatch *add.MixAndMatch
	if formData.MixAndMatch {
		mixAndMatch = &add.MixAndMatch{
			Role: formData.MixAndMatchRole,
		}
	}

	var teamCompetition *add.TeamCompetition
	if formData.TeamCompetition {
		teamCompetition = &add.TeamCompetition{
			Name: formData.TeamName,
		}
	}

	var tShirt *add.TShirt
	if formData.TShirt {
		switch add.TShirtStyle(formData.TShirtSize) {
		case add.UnisexS, add.UnisexM, add.UnisexL, add.UnisexXL, add.Unisex2XL, add.Unisex3XL, add.BellaS, add.BellaM, add.BellaL, add.BellaXL, add.Bella2XL:
			tShirt = &add.TShirt{
				Style: add.TShirtStyle(formData.TShirtSize),
			}
		case "None":
			logger.Warn("No T-shirt size submitted?")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		default:
			logger.Warnf("Could not parse tshirt style %s", formData.TShirtSize)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	var housing add.Housing
	switch formData.HousingStatus {
	case "None":
		housing = &add.NoHousing{}
	case "Require":
		housing = &add.RequireHousing{
			PetAllergies: formData.PetAllergies,
			Details:      formData.HousingRequestDetails,
		}
	case "Provide":
		housing = &add.ProvideHousing{
			Pets:     formData.MyPets,
			Quantity: formData.HousingNumber,
			Details:  formData.MyHousingDetails,
		}
	default:
		logger.Warnf("Could not parse housing status %s", formData.HousingStatus)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = addService.Add(r.Context(), &add.Registration{
		FirstName:       formData.FirstName,
		LastName:        formData.LastName,
		StreetAddress:   formData.Address,
		City:            formData.City,
		State:           formData.State,
		ZipCode:         formData.Zip,
		Email:           formData.Email,
		HomeScene:       formData.HomeScene,
		IsStudent:       formData.Student,
		PassType:        passType,
		MixAndMatch:     mixAndMatch,
		SoloJazz:        formData.SoloJazz,
		TeamCompetition: teamCompetition,
		TShirt:          tShirt,
		Housing:         housing,
	})
	if err != nil {
		logger.WithError(err).Error("Error adding regitration to backend")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "http://test.daytonswingsmackdown.com", http.StatusFound)
}
