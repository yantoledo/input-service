package http_protocol

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpService(t *testing.T) {
	key := "Content_type"
	value := []string{"application/json"}
	header := http.Header(map[string][]string{key: value})

	url := "https://www.google.com.br/"
	client := NewClientMock()
	body := NewBodyMock()

	httpService := NewHttpService()
	resp, err := httpService.Post(url, body, header, client)

	assert.Equal(t, resp, &http.Response{})
	assert.Equal(t, err, nil)
}
