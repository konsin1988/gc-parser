package etl

import (
	"context"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/parser"
	"konsin1988/gc-agent/model"
)

type SellerJob struct {
   	Services 

    SellerID string 
}


func (j *SellerJob) Fetch(ctx context.Context) (any, error) {
    return j.Ozon.Seller(ctx, j.SellerID)
}

func (j *SellerJob) Parse(data any) (any, error) {
    page := data.(*ozon.PageResponse)

    return parser.ParseSeller(page)
}

func (j *SellerJob) Save(ctx context.Context, data any) error {

    brands := data.([]model.Brand)

    for _, brand := range brands {

        brandID, err := j.Repo.InsertBrand(ctx, brand)
        if err != nil {
            return err
        }

        err = j.Repo.InsertSellerBrand(
            ctx,
            j.SellerID,
            brandID,
        )
        if err != nil {
            return err
        }
    }

    return nil
}

func NewSellerJob(
	services *Services,
	sellerID string,
) *SellerJob {
	return &SellerJob{
		Services: *services,
		SellerID: sellerID,
	}
}


func (j *SellerJob) Run(ctx context.Context) error {
	raw, err := j.Fetch(ctx)
	if err != nil {
		return err
	}

	parsed, err := j.Parse(raw)
	if err != nil {
		return err
	}

	if err = j.Save(ctx, parsed); err != nil {
		return err
	}

	return nil
}
