package http_protocol

import "net/http"

type ClientMock struct {
}

func NewClientMock() *ClientMock {
	return &ClientMock{}
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

type BodyMock struct {
}

func NewBodyMock() *BodyMock {
	return &BodyMock{}
}
