package etl

import (
	"context"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/dadata"
	"konsin1988/gc-agent/repository"
)

type Job interface {
    Fetch(ctx context.Context) (any, error)
    Parse(data any) (any, error)
    Save(ctx context.Context, data any) error
}

type Services struct {
    Ozon 			*ozon.Client
		Dadata		*dadata.Client
    Repo   		*repository.Repository
}
