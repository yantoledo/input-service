package http_protocol

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	http HttpClientProvider
}

func NewHttpClient() HttpClient {
	return HttpClient{http: &http.Client{}}
}

func (h HttpClient) Post(request HTTPRequest) (HTTPResponse, error) {
	if request.Headers == nil {
		request.Headers = make(map[string]string)
	}
	request.Headers["content-type"] = "application/json"
	return h.processRequest("POST", request)
}

func (h HttpClient) processRequest(method string, request HTTPRequest) (HTTPResponse, error) {
	httpRequest, err := http.NewRequest(method, request.URL, bytes.NewBuffer(request.Body))
	if err != nil {
		return HTTPResponse{}, err
	}

	for key, value := range request.Headers {
		httpRequest.Header.Add(key, value)
	}

	return processResponse(h.http.Do(httpRequest))
}

func processResponse(resp *http.Response, err error) (HTTPResponse, error) {
	var result HTTPResponse

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	result.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	result.StatusCode = resp.StatusCode

	return result, nil
}
