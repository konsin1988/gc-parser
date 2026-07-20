package etl

import (
	"context"
	"log"

	"konsin1988/gc-agent/repository"
	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/dadata"
	"konsin1988/gc-agent/parser"
	"konsin1988/gc-agent/model"
)

type GoodItemJob struct {
	Services	

	GoodURL			string
	QueryID			int
}


func (j *GoodItemJob) Fetch(ctx context.Context) (any, error) {
    return j.Ozon.DataByURL(ctx, j.GoodURL)
}

func (j *GoodItemJob) Parse(data any) (any, error) {
    page := data.(*ozon.PageResponse)

    return parser.ParseGoodItem(page)
}

func (j *GoodItemJob) Save(ctx context.Context, data any) error {
    good := data.(*model.GoodItem)
		err := j.Repo.InsertGoodItem(ctx, good)
		return err 
}

func NewGoodItemJob(
	ozon *ozon.Client,
	dadata  *dadata.Client,
	repo *repository.Repository,
	goodUrl string,
	queryID int,
) *GoodItemJob {
	return &GoodItemJob{
		Services: Services{
			Ozon: ozon,
			Dadata: dadata,
			Repo:   repo,
		},
		GoodURL: goodUrl,
		QueryID: queryID,
	}
}


func (j *GoodItemJob) Run(ctx context.Context) error {
	raw, err := j.Fetch(ctx)
	if err != nil {
		return err
	}

	parsed, err := j.Parse(raw)
	if err != nil {
		return err
	}

	log.Print(parsed)

	//if err = j.Save(ctx, parsed); err != nil {
	//	return err
	//}

	return nil
}
