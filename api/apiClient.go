package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiClient interface {
	Get(string, interface{}) error
}

type apiClient struct {
	host string
	port string
}

func NewApiClient(host, port string) ApiClient {
	return &apiClient{
		host: host,
		port: port,
	}
}

func (ac *apiClient) Get(route string, v interface{}) error {
	url := fmt.Sprintf("http://%s:%s/%s", ac.host, ac.port, route)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Could not GET %s. Error = %e\n", url, err)
		return err
	}
	if res.StatusCode >= 400 {
		var resp ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&resp)
		if err != nil {
			fmt.Printf("GET %s returned with error code %d. Could not parse response body.\n", url, res.StatusCode)
			return err
		}
		fmt.Printf("GET %s returned with error code %d. Message %+v\n", url, res.StatusCode, resp)
		return fmt.Errorf("GET %s returned with error code %d. Message %+v\n", url, res.StatusCode, resp)
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		fmt.Printf("Error decoding response from GET %s request. Error = %e\n", url, err)
		return err
	}
	return nil
}
