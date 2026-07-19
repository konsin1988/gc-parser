package etl

import (
	"context"
)

func Run(ctx context.Context, job Job) error {
    raw, err := job.Fetch(ctx)
    if err != nil {
        return err
    }

    parsed, err := job.Parse(raw)
    if err != nil {
        return err
    }

    return job.Save(ctx, parsed)
}

