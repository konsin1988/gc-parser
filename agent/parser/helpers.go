package parser

import (
	"fmt"
	"strings"

	"konsin1988/gc-agent/marketplace/ozon"
)

func FindWidgetKey(page *ozon.PageResponse, prefix string) (string, error) {
	for key := range page.WidgetStates {
		if strings.HasPrefix(key, prefix) {
			return key, nil
		}
	}

	return "", fmt.Errorf("widget with prefix %q not found", prefix)
}


