package modules

import (
	"errors"
	"io"
	"net/http"
)

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (c HttpClient) post(path string, value io.Reader, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, path, value)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("your apiKey is invalid")
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
