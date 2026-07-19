package dadata

import (
    "net/http"

    "konsin1988/gc-agent/config"
)

type Client struct {
    http *http.Client
    cfg  config.Config
}

func New(cfg config.Config) *Client {
    return &Client{
        http: &http.Client{},
        cfg:  cfg,
    }
}
