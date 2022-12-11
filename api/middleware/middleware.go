package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aSquidsBody/go-common/api"
	"github.com/aSquidsBody/go-common/env"
	"github.com/aSquidsBody/go-common/logs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	authv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
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

func WithIntVars(next func(http.ResponseWriter, *http.Request), names ...string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()
		for _, name := range names {
			v, ok := vars[name]
			if !ok {
				err := fmt.Errorf("Malformed URL. Missing %s", name)
				api.WriteBadRequestError(w, err)
				return
			}
			intV, err := strconv.Atoi(v)
			if err != nil {
				api.WriteBadRequestError(w, fmt.Errorf("Malformed URL. Invalid value for %s: value = %s", name, v))
			}
			ctx = context.WithValue(ctx, name, intV)
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

// Read ServiceAccount token
func Internal(next func(http.ResponseWriter, *http.Request), clientSet *kubernetes.Clientset) func(http.ResponseWriter, *http.Request) {
	serviceName := env.GetEnv("SERVICE_NAME", "")
	if serviceName == "" {
		logs.Fatal(fmt.Errorf("Env var not defined."), "SERVICE_NAME is undefined")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		clientId := r.Header.Get("X-Client-Id")
		ctx := r.Context()

		// if no clientId, then assume is external request. The
		if len(clientId) == 0 {
			ctx = context.WithValue(ctx, "internal", false)
			next(w, r.WithContext(ctx))
			return
		}

		tr := authv1.TokenReview{
			Spec: authv1.TokenReviewSpec{
				Token:     clientId,
				Audiences: []string{serviceName},
			},
		}
		result, err := clientSet.AuthenticationV1().TokenReviews().Create(context.TODO(), &tr, metav1.CreateOptions{})
		if err != nil {
			logs.Error(err, "Could not authenticate ServiceAccountToken")
			api.WriteServerError(w, fmt.Errorf("Could not authenticate ServiceAccountToken. Error = %s", err.Error()))
			return
		}
		if result.Status.Authenticated {
			ctx = context.WithValue(ctx, "internal", true)
			next(w, r.WithContext(ctx))
			return
		}
		api.WriteForbidden(w, "Invalid token")
	}
}
