package src

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/Houndie/dss-registration/src/templates"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	file, err := templates.Assets.Open("registration.html")
	if err != nil {
		http.Error(w, "could not open html document", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	registrationBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "could not open html document", http.StatusInternalServerError)
		return
	}

	t, err := template.New("registration").Parse(string(registrationBytes))
	if err != nil {
		http.Error(w, "could not open html document", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "could not open html document", http.StatusInternalServerError)
		return
	}

	return
}
