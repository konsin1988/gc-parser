package service

import (
	"net/url"
	"strings"
)


func GoodSlug(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		return ""
	}

	return parts[1]
}

