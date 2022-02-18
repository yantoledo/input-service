package http_protocol

type HTTPRequest struct {
	URL     string
	Headers map[string]string
	Body    []byte
}

type HTTPResponse struct {
	StatusCode int
	Body       []byte
}
