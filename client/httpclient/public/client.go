package public

import (
	"zb_sdk_go/client/httpclient"
)

type Client struct {
	*httpclient.HttpClient
}

func NewPublicClient(endpoint string) *Client {
	var client Client
	client.HttpClient = httpclient.NewHttpClient(endpoint)

	return &client
}
