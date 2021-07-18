package main

import (
	"errors"
	"net/http"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(2, 4)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if !limiter.Allow() {
			app.rateLimitExceededResponse(w,r)
			return
		}
		next.ServeHTTP(w, r)
	})
} 


func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.logger.PrintError(errors.New(message), nil)
}

