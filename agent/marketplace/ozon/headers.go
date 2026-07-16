package ozon

import (
    http "github.com/bogdanfinn/fhttp"
)

func (c *Client) fillHeaders(req *http.Request) {
    req.Header.Set("User-Agent", c.cfg.Headers.UserAgent)
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Accept-Language", "en-US,en;q=0.9")
    req.Header.Set("Accept-Encoding", "gzip, deflate")

    req.Header.Set("x-o3-app-name", c.cfg.Headers.AppName)
    req.Header.Set("x-o3-app-version", c.cfg.Headers.AppVersion)
    req.Header.Set("x-o3-manifest-version", c.cfg.Headers.ManifestVersion)
    req.Header.Set("x-o3-parent-requestid", c.cfg.Headers.ParentRequestID)
    req.Header.Set("x-page-view-id", c.cfg.Headers.PageViewID)
    req.Header.Set("x-page-previous", "home")

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Sec-Fetch-Dest", "empty")
    req.Header.Set("Sec-Fetch-Mode", "cors")
    req.Header.Set("Sec-Fetch-Site", "same-origin")
    req.Header.Set("Connection", "keep-alive")

    req.Header.Set("Cookie", c.cfg.Ozon.Cookie)
}
