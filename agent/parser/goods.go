package parser

import (
	"encoding/json"
	"fmt"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/model"
)

type goodsWidget struct {
	Items []goodsWidgetItem `json:"items"`
}

type goodsWidgetItem struct {
	ID     string `json:"id"`
	Action action `json:"action"`
}

type action struct {
	Link string `json:"link"`
}


func ParseGoods(page *ozon.PageResponse) (*model.GoodsPage, error) {
	gridKey, err := FindWidgetKey(page, "tileGridDesktop-")
	if err != nil {
		return nil, err
	}

	rawWidget, ok := page.WidgetStates[gridKey]
	if !ok {
		return nil, fmt.Errorf("widget %q not found", gridKey)
	}

	// WidgetStates values are JSON strings, not JSON objects.
	var widgetJSON string
	if err := json.Unmarshal(rawWidget, &widgetJSON); err != nil {
		return nil, err
	}

	var widget goodsWidget
	if err := json.Unmarshal([]byte(widgetJSON), &widget); err != nil {
		return nil, err
	}

	items := make([]model.Good, 0, len(widget.Items))

	for _, item := range widget.Items {
		items = append(items, model.Good{
			Sku:   item.ID,
			Link: item.Action.Link,
		})
	}

	return &model.GoodsPage{
		Goods:		items,
		NextPage:	page.NextPage,	
	}, nil
}
