package parser

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"regexp"

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


func parsePrice(s string) (*int, error) {
    if s == "" {
        return nil, nil
    }

    s = strings.ReplaceAll(s, "\u2009", "")
    s = strings.ReplaceAll(s, "₽", "")
    s = strings.ReplaceAll(s, " ", "")

    value, err := strconv.Atoi(s)
    if err != nil {
        return nil, err
    }

    return &value, nil
}

func ParseWidget[T any](page *ozon.PageResponse, prefix string, dst *T) error {
    key, err := FindWidgetKey(page, prefix)
    if err != nil {
        return err
    }

    raw, ok := page.WidgetStates[key]
    if !ok {
        return fmt.Errorf("widget %q not found", key)
    }

    var widgetJSON string
    if err := json.Unmarshal(raw, &widgetJSON); err != nil {
        return err
    }

    return json.Unmarshal([]byte(widgetJSON), dst)
}


func normalizeSlug(link, prefix string) string {
	link = strings.TrimPrefix(link, prefix)
	return strings.Trim(link, "/")
}

func parseBrandSlug(link string) string {
	re := regexp.MustCompile(
		`/brand/(.+?)/\?all_items=true`,
	)

	match := re.FindStringSubmatch(link)

	if len(match) < 2 {
		return ""
	}

	return match[1]
}


func parseSellerSlug(link string) string {
	parts := strings.Split(strings.Trim(link, "/"), "/")

	if len(parts) == 0 {
		return ""
	}

	return parts[len(parts)-1]
}

func parseOGRNIP(factors any) string {
	re := regexp.MustCompile(`\d{13,15}`)

	data, _ := json.Marshal(factors)

	match := re.FindString(string(data))

	return match
}


