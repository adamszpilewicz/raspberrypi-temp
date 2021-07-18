package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) listTemps(w http.ResponseWriter, r *http.Request) {
	temps, err := app.models.Temperatures.GetAllLimit()
	if err != nil {
		app.logger.PrintError(err, nil)
	}

	js, err := json.MarshalIndent(temps, "", "\t")
	if err != nil {
		app.logger.PrintError(err, nil)
	}
	w.Header().Set("Content-Type", "application/json")
	js = append(js, '\n')


	w.Write(js)

}

func (app *application) listArgs(w http.ResponseWriter, r *http.Request) {
	u, _ := r.URL.Parse(r.URL.String())
	queryParams := u.Query()
	var ls []string
	for i := range queryParams {
		i = string(i)
		ls = append(ls, i)
	}
	js, _ := json.Marshal(ls)
	w.Write(js)

}
