package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertBrand (
	ctx context.Context,
	brand model.Brand,
) (int, error) {

	var id int

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO parsing_data.brand (slug, title)
		VALUES ($1, $2)
		ON CONFLICT (slug)
		DO UPDATE SET
			title = EXCLUDED.title
		RETURNING id
	`, brand.Slug, brand.Title).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
