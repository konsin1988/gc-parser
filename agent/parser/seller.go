package parser

import (
	"encoding/json"
	"errors"
	"path"
	"strings"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/model"
)

func ParseSeller(page *ozon.PageResponse) ([]model.Brand, error) {

	type cellListWidget struct {
		Cells []struct {
			DSCell struct {
				CenterBlock struct {
					Title struct {
						Text string `json:"text"`
					} `json:"title"`
				} `json:"centerBlock"`

				Common struct {
					Action struct {
						Link string `json:"link"`
					} `json:"action"`
				} `json:"common"`
			} `json:"dsCell"`
		} `json:"cells"`
	}

	for key, rawWidget := range page.WidgetStates {

		if !strings.HasPrefix(key, "cellList-") {
			continue
		}

		// WidgetStates value is a JSON string.
		var widgetJSON string
		if err := json.Unmarshal(rawWidget, &widgetJSON); err != nil {
			continue
		}

		var widget cellListWidget
		if err := json.Unmarshal([]byte(widgetJSON), &widget); err != nil {
			continue
		}

		brands := make([]model.Brand, 0)

		for _, cell := range widget.Cells {

			link := cell.DSCell.Common.Action.Link

			// Ignore all non-brand cells.
			if !strings.Contains(link, "/brand/") {
				continue
			}

			title := cell.DSCell.CenterBlock.Title.Text
			slug := path.Base(strings.TrimRight(link, "/"))

			brands = append(brands, model.Brand{
				Title: title,
				Slug:  slug,
			})
		}

		// This is the widget we need.
		if len(brands) > 0 {
			return brands, nil
		}
	}

	return nil, errors.New("no brands found")
}
