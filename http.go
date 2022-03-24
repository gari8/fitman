package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func post(path string, value io.Reader, header map[string]string) ([]byte, error) {
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
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
