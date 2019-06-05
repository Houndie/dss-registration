package api

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/Houndie/dss-registration/templates"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	file, err := templates.Assets.Open("registration.html")
	if err != nil {
		http.Error(w, "Could not open html file", 500)
		return
	}
	defer file.Close()
	registrationBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Could not read html file", 500)
		return
	}
	t, err := template.New("registration").Parse(string(registrationBytes))
	if err != nil {
		http.Error(w, "Could not parse html file", 500)
		return
	}
	t.Execute(w, nil)
}
