package api

import (
	"io"
	"log"
	"net/http"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func (api *API) RetrieveTodo() ([]byte, error) {
	req, err := http.NewRequest("GET", api.BaseURL, nil)
	if err != nil {
		log.Default().Fatal("Error creating request")
		return nil, err
	}
	resp, err := api.Client.Do(req)
	if err != nil {
		//log.Default().Fatal("Error sending request")
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Default().Fatal("Error closing body")
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Default().Fatal("Error reading body")
		return nil, err
	}

	return body, nil
}
