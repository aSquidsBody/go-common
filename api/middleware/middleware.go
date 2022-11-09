package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aSquidsBody/go-common/api"
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
