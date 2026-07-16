package etl

import (
	"context"

	"konsin1988/gc-agent/repository"
	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/parser"
	_ "konsin1988/gc-agent/model"
)

type SearchGoodsJob struct {
	BaseJob 

  SearchURL 	string
	CategoryId  int
	QueryId			int
	maxPages 		int
}


func (j *SearchGoodsJob) Fetch(ctx context.Context) (any, error) {
    return j.Client.GoodsBySearch(ctx, j.SearchURL)
}

func (j *SearchGoodsJob) Parse(data any) (any, error) {
    page := data.(*ozon.PageResponse)

    return parser.ParseGoods(page)
}

func (j *SearchGoodsJob) Save(ctx context.Context, data any) error {
    goods := data.(*parser.GoodsPage)

    return j.Repo.SaveGoods(ctx, j.CategoryId, j.QueryId, goods.Goods)
}

func NewSearchGoodsJob(
	client *ozon.Client,
	repo *repository.Repository,
	searchText string,
	categoryID int,
	queryID int,
	maxPages int,
) *SearchGoodsJob {
	return &SearchGoodsJob{
		BaseJob: BaseJob{
			Client: client,
			Repo:   repo,
		},
		SearchURL: client.BuildSearchPageURL(searchText),
		CategoryId: categoryID,
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

		if err := j.Repo.SaveGoods(ctx, j.CategoryId, j.QueryId, parsed.Goods); err != nil {
			return err
		}

		if parsed.NextPage == "" {
			break
		}

		j.SearchURL = parsed.NextPage
	}

	return nil
}
