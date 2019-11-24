package dynamic

import (
	"encoding/json"
	"net/http"
)

func PopulateForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		logger.Info("Populate Form (CORS Preflight)")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	logger.Info("Populate Form")
	res, err := populateService.Populate(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	httpres := struct {
		WeekendPassCost int `json:"weekend_pass_cost"`
		WeekendPassTier int `json:"weekend_pass_tier"`
		DancePassCost   int `json:"dance_pass_cost"`
		MixAndMatchCost int `json:"mix_and_match_cost"`
		SoloJazzCost    int `json:"solo_jazz_cost"`
		TeamCompCost    int `json:"team_comp_cost"`
		TShirtCost      int `json:"tshirt_cost"`
		StudentDiscount int `json:"student_discount"`
	}{
		WeekendPassCost: res.WeekendPassCost,
		WeekendPassTier: res.WeekendPassTier,
		DancePassCost:   res.DancePassCost,
		MixAndMatchCost: res.MixAndMatchCost,
		SoloJazzCost:    res.SoloJazzCost,
		TeamCompCost:    res.TeamCompCost,
		TShirtCost:      res.TShirtCost,
		StudentDiscount: res.StudentDiscount,
	}
	bytes, err := json.Marshal(&httpres)
	if err != nil {
		logger.WithError(err).Error("Error converting populate form response to json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		logger.WithError(err).Error("Error writing response body")
	}
}
