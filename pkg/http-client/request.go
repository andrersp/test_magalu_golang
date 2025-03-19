package httpclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Request interface {
	SetResult(output any) Request
	Get(url string) (*http.Response, error)
	Bytes() []byte
}

type request struct {
	ctx           context.Context
	client        *client
	result        any
	bytesResponse io.ReadCloser
}

// Get implements Request.
func (r *request) Get(url string) (*http.Response, error) {
	resp, err := r.client.execute(r.ctx, http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	r.bytesResponse = resp.Body

	return resp, r.parseResult()
}

// SetResult implements Request.
func (r *request) SetResult(output any) Request {
	r.result = output

	return r
}

func (r *request) parseResult() error {
	if r.result == nil {
		return nil
	}

	bts := r.Bytes()

	return json.Unmarshal(bts, r.result)
}

// Bytes implements Request.
func (r *request) Bytes() []byte {
	bts, _ := io.ReadAll(r.bytesResponse)
	return bts
}
