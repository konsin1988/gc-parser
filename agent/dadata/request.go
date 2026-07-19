package dadata

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func (c *Client) newRequest(body any) (*http.Request, error) {

    b, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(
        "POST",
        c.cfg.Dadata.DadataApiURL,
        bytes.NewReader(b),
    )
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set(
        "Authorization",
        "Token "+c.cfg.Dadata.DadataApiToken,
    )

    return req, nil
}
