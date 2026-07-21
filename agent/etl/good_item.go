package etl

import (
	"context"
	"log"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/parser"
	"konsin1988/gc-agent/model"
	"konsin1988/gc-agent/service"
)

type GoodItemJob struct {
	*Services	

	GoodItemService 	*service.GoodItemService
	GoodURL						string
	QueryID						int
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
	services *Services,
	goodItemService  *service.GoodItemService,
	goodUrl string,
	queryID int,
) *GoodItemJob {
	return &GoodItemJob{
		Services: services,
		GoodItemService: goodItemService,
		GoodURL: goodUrl,
		QueryID: queryID,
	}
}


func (j *GoodItemJob) Run(ctx context.Context) error {
	raw, err := j.Fetch(ctx)
	if err != nil {
		return err
	}

	parsedRaw, err := j.Parse(raw)
	if err != nil {
		return err
	}
	parsed := parsedRaw.(*model.ParsedGoodItem)

	log.Print(parsed)
	log.Print("==========================")
	if err := j.GoodItemService.ProcessGoodItem(ctx, parsed, j.GoodURL, j.QueryID); err != nil {
    	return err
  }

  if parsed.Seller != nil && parsed.Seller.ID != "0" && parsed.Seller.ID != "" {
    sellerJob := SellerJob{
      Services: j.Services,
      SellerID: parsed.Seller.ID,
    }
    if err := sellerJob.Run(ctx); err != nil {
        return err
    }
  }

  if parsed.ReviewLink != "" {
    reviewJob := ReviewJob{
        Services:  j.Services,
        ReviewURL: parsed.ReviewLink,
				MaxPages: 5,
    }
    if err := reviewJob.Run(ctx); err != nil {
        return err
    }
  }

	return nil
}
