package ozon

import (
  tls_client "github.com/bogdanfinn/tls-client"
  "github.com/bogdanfinn/tls-client/profiles"

	"konsin1988/gc-agent/config"
)

type Client struct {
    http 	tls_client.HttpClient
    cfg 	config.Config 
}

func New(cfg config.Config) (*Client, error) {
    jar := tls_client.NewCookieJar()

    httpClient, err := tls_client.NewHttpClient(
        tls_client.NewNoopLogger(),
        tls_client.WithClientProfile(profiles.Firefox_133),
        tls_client.WithCookieJar(jar),
    )
    if err != nil {
        return nil, err
    }

    return &Client{
        http: httpClient,
        cfg:  cfg,
    }, nil
}
