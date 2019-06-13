package webapi

import (
	"context"
	"fmt"
	"net/http"
)

const url = "https://qiita.com/api/v2"

func createQiitaClient() (*Client, error) {
	endpointUrl := "https://qiita.com/api/v2"
	httpClient := &http.Client{}

	client, err := newClient(endpointUrl, httpClient)
	return client, err
}

func Items(ctx context.Context, pathPram string) error {

	c, err := createQiitaClient()
	if err != nil {
		fmt.Println(err)
		return err
	}
	req, err := c.newRequest(ctx, "GET", "/items", nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	return nil
}
