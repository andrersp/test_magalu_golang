package httpclient

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type HttpClient interface {
	NewRequest(ctx context.Context) Request
	SetOutputDir(dir string) HttpClient
	SetRetryCount(count int) HttpClient
	execute(ctx context.Context, method, url string, payload io.Reader) (*http.Response, error)
}

type client struct {
	httpClient *http.Client
	outputDir  string
	retryCount int
}

func NewHttpClient() HttpClient {
	httpClient := http.Client{}

	client := new(client)
	client.httpClient = &httpClient
	client.retryCount = 3

	return client
}

// SetRetryCount implements HttpClient.
func (c *client) SetRetryCount(count int) HttpClient {
	c.retryCount = count
	return c
}

// SetOutputDir implements HttpClient.
func (c *client) SetOutputDir(dir string) HttpClient {
	c.outputDir = dir

	return c
}

func (c *client) NewRequest(ctx context.Context) Request {
	req := new(request)
	req.ctx = ctx
	req.client = c

	return req
}

func (c *client) execute(ctx context.Context,
	method string, url string, payload io.Reader) (response *http.Response, err error) {
	if c.retryCount == 1 {
		return c.makeRequest(ctx, method, url, payload)
	}

	for index := 0; index < c.retryCount; index++ {
		response, err = c.makeRequest(ctx, method, url, payload)
		if err == nil {
			break
		}

		slog.Info("httpClient",
			"message", fmt.Sprintf("error on attempts %d of %d", index+1, c.retryCount),
			"err", err,
		)
	}

	return
}

func (c *client) makeRequest(ctx context.Context, method string,
	url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}
