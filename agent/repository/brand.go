package repository

import (
	"context"
	"database/sql"
	"errors"

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


func (r *Repository) GetBrandByID (
	ctx context.Context,
	brandID	int,
) (*model.Brand, error) {
	var brand model.Brand
	err := r.db.QueryRowContext(ctx, `
        SELECT slug, title
        FROM parsing_data.brand
        WHERE id = $1
    `, brandID).Scan(
        &brand.Slug,
        &brand.Title,
    )

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil 
        }
        return nil, err
    }
    return &brand, nil

}
