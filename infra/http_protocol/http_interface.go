package http_protocol

import "net/http"

type HttpClientProvider interface {
	Do(request *http.Request) (*http.Response, error)
}
