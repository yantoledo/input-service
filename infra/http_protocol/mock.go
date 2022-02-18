package http_protocol

import "net/http"

type ClientProviderMock struct {
}

func NewClientProviderMock() *ClientProviderMock {
	return &ClientProviderMock{}
}

func (c *ClientProviderMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

type HttpClientMock struct {
	http ClientProviderMock
}

func NewHttpClientMock() *HttpClientMock {
	return &HttpClientMock{http: ClientProviderMock{}}
}

func (c *HttpClientMock) Post(request HTTPRequest) (HTTPResponse, error) {
	return HTTPResponse{}, nil
}

type BodyMock struct {
}

func NewBodyMock() *BodyMock {
	return &BodyMock{}
}
