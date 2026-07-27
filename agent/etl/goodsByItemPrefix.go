package etl

import (
	"context"
	"log"
	"sync"
	"fmt"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/parser"
	"konsin1988/gc-agent/service"
	"konsin1988/gc-agent/model"
)

type GoodsByItemPrefixJob struct {
	*Services	

	ItemID			int
	SearchBase	string
	SearchURL		string
	maxPages 		int
}


func (j *GoodsByItemPrefixJob) Fetch(ctx context.Context) (any, error) {
    return j.Ozon.DataByURL(ctx, j.SearchURL)
}

func (j *GoodsByItemPrefixJob) Parse(data any) (any, error) {
    page := data.(*ozon.PageResponse)

    return parser.ParseGoods(page)
}

func (j *GoodsByItemPrefixJob) Save(ctx context.Context, data any) error {
		return nil
}

func NewGoodsByItemPrefixJob(
	services *Services,
	searchBase  string,
	itemID  int,
	maxPages int,
) *GoodsByItemPrefixJob {
	return &GoodsByItemPrefixJob{
		Services: services,
		SearchBase: searchBase,
		ItemID: itemID,
		maxPages: maxPages,
	}
}


func (j *GoodsByItemPrefixJob) getSlugByID (ctx context.Context ) (string, error) {
	switch j.SearchBase {
	case "brand":
		brand, err := j.Services.Repo.GetBrandByID(
			ctx,
			j.ItemID,
		)
		if err != nil {
			return "", err 
		}
		return brand.Slug, nil

	case "category":
		category, err := j.Services.Repo.GetCategoryByID(
			ctx,
			j.ItemID,
		)
		if err != nil {
			return "", err 
		}
		return category.Slug, nil

	case "seller":
		seller, err := j.Services.Repo.GetSellerByID(
			ctx,
			j.ItemID,
		)
		if err != nil {
			return "", err 
		}
		return seller.Slug, nil
	}

	return "", nil
}


func (j *GoodsByItemPrefixJob) Run(ctx context.Context) error {
	goodItemService := service.NewGoodItemService(
		j.Services.Repo,
		j.Services.Ozon,
		j.Services.Dadata,
	)

	slug, err := j.getSlugByID(ctx) 
	if err != nil {
		return err
	}

	j.SearchURL = fmt.Sprintf("/%s/%s/", j.SearchBase, slug)

	log.Print(j.SearchURL)

	queryID, err := j.Services.Repo.InsertQuery(
		ctx,
		j.SearchURL,
	)
	if err != nil {
		log.Fatal(err)
	}


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
						queryID,
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
