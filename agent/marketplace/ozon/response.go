package ozon

import (
	"encoding/json"
)

type PageResponse struct {
	PageInfo struct {
		AnalyticsInfo struct {
			Sku json.Number `json:"sku"`
		} `json:"analyticsInfo"`
	} `json:"pageInfo"`

	WidgetStates 	map[string]json.RawMessage 	`json:"widgetStates"`
	NextPage			string 											`json:"nextPage,omitempty"`
}


func DecodePageResponse(body []byte) (*PageResponse, error) {
	var page PageResponse

	if err := json.Unmarshal(body, &page); err != nil {
		return nil, err
	}

	return &page, nil
}
