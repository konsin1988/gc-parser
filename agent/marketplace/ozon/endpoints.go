package ozon

import  (
	"context"
	"net/url"
)

func (c *Client) BuildSearchPageURL(
    searchText string,
) string {

    u := &url.URL{
        Path: "/search/",
    }

    q := u.Query()
    q.Set("from_global", "true")
    q.Set("layout_container", "default")
    q.Set("layout_page_index", "3")
    q.Set("page", "3")
    q.Set("paginator_token", "3635012")
    q.Set("text", searchText)

    u.RawQuery = q.Encode()

    return u.String()
}

func (c *Client) GoodsBySearch(ctx context.Context, url string) (*PageResponse, error) {
    req, err := c.newPageRequest(url)
    if err != nil {
        return nil, err
    }

    req = req.WithContext(ctx)

		body, err := c.do(req)
		if err != nil {
			return nil, err
		}

		return DecodePageResponse(body)
}
