package etl

import (
	"context"
	"log"
	"sync"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/parser"
	"konsin1988/gc-agent/service"
	"konsin1988/gc-agent/model"
)

type SearchGoodsJob struct {
	*Services	

  SearchURL 	string
	QueryId			int
	maxPages 		int
}


func (j *SearchGoodsJob) Fetch(ctx context.Context) (any, error) {
    return j.Ozon.DataByURL(ctx, j.SearchURL)
}

func (j *SearchGoodsJob) Parse(data any) (any, error) {
    page := data.(*ozon.PageResponse)

    return parser.ParseGoods(page)
}

func (j *SearchGoodsJob) Save(ctx context.Context, data any) error {
		return nil
}

func NewSearchGoodsJob(
	services *Services,
	searchText string,
	queryID int,
	maxPages int,
) *SearchGoodsJob {
	return &SearchGoodsJob{
		Services: services,
		SearchURL: services.Ozon.BuildSearchPageURL(searchText),
		QueryId: queryID,
		maxPages: maxPages,
	}
}


func (j *SearchGoodsJob) Run(ctx context.Context) error {
	goodItemService := service.NewGoodItemService(
		j.Services.Repo,
		j.Services.Ozon,
		j.Services.Dadata,
	)

	for i := 0; i < j.maxPages; i++ {
		raw, err := j.Fetch(ctx)
		if err != nil {
			return err
		}
		page := raw.(*ozon.PageResponse)
		parsed, err := parser.ParseGoods(page)
		if err != nil {
			return err
		}


		var wg sync.WaitGroup

		for _, good := range parsed.Goods {
			sem := make(chan struct{}, 3) // max 10 concurrent jobs

			wg.Add(1)
	    go func(g model.Good) {
					defer wg.Done()

					sem <- struct{}{}
    			defer func() { <-sem }()
					goodItemJob := NewGoodItemJob(
						j.Services,
						goodItemService,
						g.Link,
						j.QueryId,
					)

					if err := goodItemJob.Run(ctx); err != nil {
						log.Print(err)
					}
	    }(good)
		}

		wg.Wait()

		if parsed.NextPage == "" {
			break
		}
		j.SearchURL = parsed.NextPage
	}
	return nil
}
