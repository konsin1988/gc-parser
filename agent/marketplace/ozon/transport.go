package ozon

import (
	"io"
	http "github.com/bogdanfinn/fhttp"
)
func (c *Client) do(req *http.Request) ([]byte, error) {
    resp, err := c.http.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}
