package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertGoodItem (
	ctx context.Context,
	goodItem *model.GoodItem,
) (error) {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO parsing_data.good_item (
																		sku, 
																		slug,
																		title, 
																		price, 
																		card_price, 
																		original_price, 
																		availability,
																		seller_id,
																		brand_id,
																		review_link
																	)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (sku)
		DO UPDATE SET
			slug = EXCLUDED.slug,
			title = EXCLUDED.title,
			price = EXCLUDED.price,
			card_price = EXCLUDED.card_price,
			original_price = EXCLUDED.original_price,
			availability = EXCLUDED.availability,
			seller_id = EXCLUDED.seller_id,
			brand_id = EXCLUDED.brand_id;
	`, goodItem.Sku, 
		goodItem.Slug,
		goodItem.Title,
		*goodItem.Price,
		*goodItem.CardPrice,
		*goodItem.OriginalPrice,
		goodItem.Availability,
		*goodItem.SellerId,
		*goodItem.BrandId,
		goodItem.ReviewLink,
	)

	return err 
}
