package fakeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type transport struct {
	client http.Client
}

func newTransport() *transport {
	return &transport{
		client: http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func (t *transport) call(url, raw, method string, body interface{}) ([]byte, int, error) {
	switch {
	case method == http.MethodGet || method == http.MethodDelete:
		return t.get(url, raw, method)
	case method == http.MethodPost || method == http.MethodPut:
		return t.post(url, raw, method, body)
	default:
		return nil, 0, fmt.Errorf("no found method: %s", method)
	}
}

func (t *transport) get(url, raw, method string) ([]byte, int, error) {
	req, err := http.NewRequest(method, url+raw, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func (t *transport) post(url, raw, method string, in interface{}) ([]byte, int, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest(method, url+raw, bytes.NewBuffer(b))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
