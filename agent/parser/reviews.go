package parser

import (
	"encoding/json"
	"fmt"
	"time"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/model"
)

type reviewWidget struct {
	Reviews []struct {
		UUID      string `json:"uuid"`
		CreatedAt int64  `json:"createdAt"`
		ItemID    json.Number `json:"itemId"`

		Author struct {
			GUID string `json:"guid"`
		} `json:"author"`

		Content struct {
			Score 	int     `json:"score"`
			Comment  string `json:"comment"`
			Positive string `json:"positive"`
			Negative string `json:"negative"`

			Photos []struct {
				URL string `json:"url"`
			} `json:"photos"`
		} `json:"content"`
	} `json:"reviews"`
}

type ReviewPage struct {
	Reviews  []model.Review
	NextPage string
}

func ParseReview(page *ozon.PageResponse) (*ReviewPage, error) {

	reviewKey, err := FindWidgetKey(page, "webListReviews-")
	if err != nil {
		return nil, err
	}

	rawWidget, ok := page.WidgetStates[reviewKey]
	if !ok {
		return nil, fmt.Errorf("widget %q not found", reviewKey)
	}

	// WidgetStates values are JSON strings.
	var widgetJSON string
	if err := json.Unmarshal(rawWidget, &widgetJSON); err != nil {
		return nil, err
	}

	var widget reviewWidget
	if err := json.Unmarshal([]byte(widgetJSON), &widget); err != nil {
		return nil, err
	}

	reviews := make([]model.Review, 0, len(widget.Reviews))

	for _, r := range widget.Reviews {


		review := model.Review{
			UUID:       r.UUID,
			CreatedAt:  time.Unix(r.CreatedAt, 0),
			Sku:        r.ItemID.String(),
			AuthorGuid: r.Author.GUID,
			Score:			r.Content.Score,
			Comment:    r.Content.Comment,
			Positive:   r.Content.Positive,
			Negative:   r.Content.Negative,
			ReviewImages: make([]model.ReviewImage, 0, len(r.Content.Photos)),
		}

		for _, photo := range r.Content.Photos {
			review.ReviewImages = append(review.ReviewImages, model.ReviewImage{
				ReviewUUID: r.UUID,
				ImgURL:     photo.URL,
			})
		}

		reviews = append(reviews, review)
	}

	return &ReviewPage{
		Reviews:  reviews,
		NextPage: page.NextPage,
	}, nil
}
