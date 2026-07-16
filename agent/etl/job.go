package etl

import (
	"context"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/repository"
)

type Job interface {
    Fetch(ctx context.Context) (any, error)
    Parse(data any) (any, error)
    Save(ctx context.Context, data any) error
}

type BaseJob struct {
    Client *ozon.Client
    Repo   *repository.Repository
}
