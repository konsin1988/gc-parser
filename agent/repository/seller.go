package repository

import (
	"context"
	"database/sql"
	"errors"

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


func (r *Repository) GetSellerByID (
	ctx context.Context,
	sellerID	int,
) (*model.Seller, error) {
	var seller model.Seller
	err := r.db.QueryRowContext(ctx, `
        SELECT slug
        FROM parsing_data.seller
        WHERE id = $1
    `, sellerID).Scan(
        &seller.Slug,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil 
        }
        return nil, err
    }
    return &seller, nil
}
