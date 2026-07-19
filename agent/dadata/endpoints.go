package dadata

import (
    "context"
    "encoding/json"
    "fmt"
)

func (c *Client) FindByOGRNIP(
    ctx context.Context,
    ogrnip string,
) (*Company, error) {

    req, err := c.newRequest(map[string]string{
        "query": ogrnip,
    })
    if err != nil {
        return nil, err
    }

    req = req.WithContext(ctx)

    resp, err := c.http.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result FindByIDResponse

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    if len(result.Suggestions) == 0 {
        return nil, fmt.Errorf("company not found")
    }

		s := result.Suggestions[0]
		
		return &Company{
		    Name: s.Value,
		    INN:  s.Data.INN,
		}, nil
}
