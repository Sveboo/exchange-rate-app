package currencyservice

import "net/http"

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type HTTPClient struct {
	client http.Client
}

func NewHTTPClient(client http.Client) HTTPClient {
	return HTTPClient{
		client: client,
	}
}

func (hc HTTPClient) Do(req *http.Request) (*http.Response, error) {
	return hc.client.Do(req)
}
