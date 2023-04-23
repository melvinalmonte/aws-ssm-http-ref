package api

import (
	"go.uber.org/zap"
	"io"
	"net/http"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func (api *API) RetrieveTodo() ([]byte, error) {
	zap.S().Info("Retrieving todo")
	req, err := http.NewRequest("GET", api.BaseURL, nil)
	if err != nil {
		zap.S().Error("Error creating request")
		return nil, err
	}
	zap.S().Info("Executing request")
	resp, err := api.Client.Do(req)
	if err != nil {
		zap.S().Error("Error executing request")
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		zap.S().Info("Closing body")
		err := Body.Close()
		if err != nil {
			zap.S().Error("Error closing body")
			return
		}
	}(resp.Body)

	zap.S().Info("Reading body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.S().Error("Error reading body")
		return nil, err
	}

	zap.S().Info("Returning request body")
	return body, nil
}
