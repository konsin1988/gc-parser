package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertSeller(
	ctx context.Context,
	seller model.Seller,
) (error) {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO parsing_data.seller (id, name, slug, ogrn_ogrnip, inn )
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id)
		DO UPDATE SET
		    name = EXCLUDED.name,
		    slug = EXCLUDED.slug,
		    ogrn_ogrnip = EXCLUDED.ogrn_ogrnip,
		    inn = EXCLUDED.inn;
	`, seller.ID, seller.Name, seller.Slug, seller.Ogrn, seller.Inn)

	return err 
}
