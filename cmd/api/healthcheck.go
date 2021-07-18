package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["api version"] = version
	data["status"] = "available"
	data["env"] = app.config.env

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "error while preparing json response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	js = append(js, '\n')

	w.Write(js)

}
