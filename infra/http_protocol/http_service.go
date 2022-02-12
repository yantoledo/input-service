package http_protocol

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpService struct {
}

func NewHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) Post(url string, body interface{}, headers http.Header, client HttpClient) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	// client := &http.Client{}
	return client.Do(request)
}
