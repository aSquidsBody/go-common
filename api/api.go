package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/aSquidsBody/go-common/logs"
)

var serviceAccountTokens map[string]string

func UseServiceAccount(refreshRate time.Duration, tokenFiles ...string) func() {
	serviceAccountTokens = make(map[string]string)

	readTokens(tokenFiles...)

	ticker := time.NewTicker(refreshRate)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				readTokens(tokenFiles...)
			}
		}
	}()

	return func() {
		close(done)
	}
}

func readTokens(tokenFiles ...string) {
	for _, tokenFile := range tokenFiles {
		b, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			panic(err)
		}

		basename := filepath.Base(tokenFile)
		serviceAccountTokens[basename] = string(b)
	}
	logs.Info("Refreshing service account token")
}

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type internalClient struct {
	*http.Client
	service string
}

func (ic *internalClient) Do(r *http.Request) (*http.Response, error) {
	token, ok := serviceAccountTokens[ic.service]
	if !ok {
		logs.Fatal("Could not make internal HTTP request to service "+ic.service, fmt.Errorf("Token not in memory"))
	}
	r.Header.Add("X-Client-Id", token)
	return ic.Client.Do(r)
}

func NewInternalClient(service string) Client {
	return &internalClient{
		Client:  &http.Client{},
		service: service,
	}
}
