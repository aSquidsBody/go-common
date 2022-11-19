package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aSquidsBody/go-common/api"
	"github.com/aSquidsBody/go-common/env"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func WithVars(next func(http.ResponseWriter, *http.Request), urlVars ...string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()
		for _, urlVar := range urlVars {
			v, ok := vars[urlVar]
			if !ok {
				err := fmt.Errorf("Malformed URL. Missing %s", urlVar)
				api.WriteBadRequestError(w, err)
				return
			}
			ctx = context.WithValue(ctx, urlVar, v)
		}
		next(w, r.WithContext(ctx))
	}
}

func WithCors(r *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{
		"Content-Disposition", "Pragma", "Accept", "Accept-Language", "Content-Type", "Accept-Encoding", "Cache-Control", "User-Agent",
		"Access-Control-Request-Method", "Connection", "Referer", "Sec-Fetch-Mode", "Access-Control-Request-Headers"})
	originsOk := handlers.AllowedOrigins([]string{env.GetEnv("ORIGIN_ALLOWED", "*")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(headersOk, originsOk, methodsOk)(r)
}
