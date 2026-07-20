package etl

import (
	"context"
	_ "log"

	"konsin1988/gc-agent/repository"
	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/dadata"
	"konsin1988/gc-agent/parser"
	_ "konsin1988/gc-agent/model"
)

type SearchGoodsJob struct {
	Services	

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
	ozon *ozon.Client,
	dadata  *dadata.Client,
	repo *repository.Repository,
	searchText string,
	queryID int,
	maxPages int,
) *SearchGoodsJob {
	return &SearchGoodsJob{
		Services: Services{
			Ozon: ozon,
			Dadata: dadata,
			Repo:   repo,
		},
		SearchURL: ozon.BuildSearchPageURL(searchText),
		QueryId: queryID,
		maxPages: maxPages,
	}
}


func (j *SearchGoodsJob) Run(ctx context.Context) error {
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

		//for _, good := range goodsPage.Goods {
	  //  go func(g model.GoodLink) {
	  //      job := &GoodItemJob{...}
	  //      _ = job.Run(ctx)
	  //  }(good)
		//}

		if parsed.NextPage == "" {
			break
		}
		j.SearchURL = parsed.NextPage
	}
	return nil
}
