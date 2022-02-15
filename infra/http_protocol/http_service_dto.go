package http_protocol

// HTTPRequest ...
type HTTPRequest struct {
	URL     string
	Headers map[string]string
	Body    []byte
}

// HTTPResponse ...
type HTTPResponse struct {
	StatusCode int
	Body       []byte
}
