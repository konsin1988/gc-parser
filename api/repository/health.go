package repository 

import "context"

func (r *Repository) Ping(ctx context.Context) error {
    return r.db.PingContext(ctx)
}
