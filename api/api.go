package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
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
	Get(string, interface{}) error
}

type internalClient struct {
	*http.Client
	service string
	url     string
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

func (ic *internalClient) Get(route string, v interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", ic.url, route), nil)
	if err != nil {
		logs.Error(fmt.Sprintf("Could not create GET request for %s", ic.service), err)
		return err
	}

	res, err := ic.Do(req)
	if err != nil {
		logs.Error(fmt.Sprintf("Could not make GET request for %s", ic.service), err)
		return err
	}

	if res.StatusCode >= 400 {
		var resp ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&resp)
		if err != nil {
			logs.Error(fmt.Sprintf("GET request to %s failed with a %d response. Could not decode response body", ic.service, res.StatusCode), err)
			return err
		}
		logs.Error(fmt.Sprintf("GET request to %s failed with a %d response. Reason = %+v", ic.service, res.StatusCode, resp), fmt.Errorf("No error"))
		return fmt.Errorf("Failed GET request. Response code %d. Response %+v", res.StatusCode, resp)
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		logs.Error(fmt.Sprintf("Could not decode response from GET request to %s", ic.service), err)
		return err
	}
	return nil
}

func NewInternalClient(host, port string) Client {
	strs := strings.Split(host, ".")
	if len(strs) <= 1 {
		logs.Fatal("Could not determine service name for internal client", fmt.Errorf("Unable to split host %s", host))
	}

	return &internalClient{
		Client:  &http.Client{},
		service: strs[0],
		url:     fmt.Sprintf("http://%s:%s", host, port),
	}
}
