package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheck)
	router.HandlerFunc(http.MethodGet, "/v1/temps", app.listTemps)
	router.HandlerFunc(http.MethodGet, "/v1/temps/args", app.listArgs)

	return app.rateLimit(router)
}
