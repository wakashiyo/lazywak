package webapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

func newClient(endpointURL string, httpClient *http.Client) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, err
	}
	client := &Client{
		URL:        parsedURL,
		HTTPClient: httpClient,
	}
	return client, nil
}

func (c *Client) newRequest(ctx context.Context, method string, subPath string, body io.Reader) (*http.Request, error) {
	endpointURL := *c.URL
	endpointURL.Path = path.Join(c.URL.Path, subPath)

	req, err := http.NewRequest(method, endpointURL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *Client) call(ctx context.Context, pathParam string) error {
	req, err := c.newRequest(ctx, "GET", "/items", nil)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
