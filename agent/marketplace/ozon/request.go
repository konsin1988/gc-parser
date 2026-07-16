package ozon

import (
	"net/url"

	http "github.com/bogdanfinn/fhttp"
)

func (c *Client) newPageRequest(pageURL string) (*http.Request, error) {
		u, err := url.Parse(c.cfg.Ozon.APIBaseURL)
    if err != nil {
        return nil, err
    }
		q := u.Query()
    q.Set("page_changed", "true")
    q.Set("url", pageURL)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
    if err != nil {
        return nil, err
    }

    c.fillHeaders(req)
    return req, nil
}

