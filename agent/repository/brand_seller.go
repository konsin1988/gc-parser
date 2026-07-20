package repository

import (
	"context"
)

func (r *Repository) InsertSellerBrand(
	ctx context.Context,
	seller_id string,
	brand_id int,
) (error) {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO parsing_data.brand_seller (seller_id, brand_id)
		VALUES ($1, $2)
		ON CONFLICT (seller_id, brand_id)
		DO NOTHING;
	`, seller_id, brand_id)

	return err 
}
