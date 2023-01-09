package fakeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slices"
)

var errStatusCode = "error, required: %v, status_code: %d"

type Client struct {
	url  string
	http *http.Client
}

func New(url string, http *http.Client) *Client {
	return &Client{url: url, http: http}
}

func (c *Client) get(url string, out interface{}, needStatusCode ...int) error {
	return c.call(url, http.MethodGet, "", nil, out, needStatusCode...)
}

func (c *Client) delete(url string, out interface{}, needStatusCode ...int) error {
	return c.call(url, http.MethodDelete, "", nil, out, needStatusCode...)
}

func (c *Client) post(url string, in, out interface{}, needStatusCode ...int) error {
	return c.call(url, http.MethodPost, "application/json", in, out, needStatusCode...)
}

func (c *Client) put(url string, in, out interface{}, needStatusCode ...int) error {
	return c.call(url, http.MethodPut, "application/json", in, out, needStatusCode...)
}

func (c *Client) call(url, method, contentType string, in, out interface{}, needStatusCode ...int) (err error) {
	var body []byte

	if in != nil {
		body, err = json.Marshal(in)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	if !slices.Contains(needStatusCode, resp.StatusCode) {
		return fmt.Errorf(errStatusCode, needStatusCode, resp.StatusCode)
	}

	if out != nil {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(body, out); err != nil {
			return err
		}
	}

	return nil
}
