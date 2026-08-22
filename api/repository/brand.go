package repository

import (
    "context"
		_ "fmt"
		_ "strings"
		_ "errors"
		_ "database/sql"

		_ "github.com/lib/pq"

		"konsin1988/gc-api/model"
)


// #################################################### ALL BRANDS 
func (r *Repository) AllBrands(
    ctx context.Context,
) ([]model.AllBrandsItem, error) {

    query := `
        SELECT
            b.id,
            b.title,
            b.slug
        FROM parsing_data.brand b
				ORDER BY b.title; 
    `

    rows, err := r.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    brands := make([]model.AllBrandsItem, 0)

    for rows.Next() {
        var brand model.AllBrandsItem

        if err := rows.Scan(
            &brand.ID,
            &brand.Title,
            &brand.Slug,
        ); err != nil {
            return nil, err
    		}

        brands = append(brands, brand)
		}
		if err := rows.Err(); err != nil {
        return nil, err
    }

		return brands, err
}
