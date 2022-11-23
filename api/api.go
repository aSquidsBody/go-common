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
var tokenFiles []string

func UseServiceAccount(refreshRate time.Duration, files ...string) func() {
	tokenFiles = files
	serviceAccountTokens = make(map[string]string)

	readTokens()

	ticker := time.NewTicker(refreshRate)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				readTokens()
			}
		}
	}()

	return func() {
		close(done)
	}
}

func readTokens() {
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

	var firstTry http.Request = *r
	var secondTry http.Request = *r

	firstTry.Header.Add("X-Client-Id", token)
	res, err := ic.Client.Do(&firstTry)
	if err != nil {
		return res, err
	}

	if res.StatusCode == http.StatusForbidden {
		// trigger a new refresh
		readTokens()
		secondTry.Header.Add("X-Client-Id", token)
		res, err = ic.Client.Do(&secondTry)
	}
	return res, err
}

func NewInternalClient(service string) Client {
	return &internalClient{
		Client:  &http.Client{},
		service: service,
	}
}
