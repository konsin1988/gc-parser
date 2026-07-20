package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertGood (
	ctx context.Context,
	catID int,
	queryID int,
	good model.Good,
) (error) {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO parsing_data.good (sku, cat_id, query_id, glink)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (sku, cat_id, query_id)
		DO UPDATE SET
			glink = EXCLUDED.glink
	`, good.Sku, catID, queryID, good.Link)

	return err 
}
