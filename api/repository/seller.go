package repository

import (
    "context"
		"fmt"
		"strings"

		"github.com/lib/pq"

		"konsin1988/gc-api/model"
)


// ##################################################### SELLER LIST
func (r *Repository) SellerList(
    ctx context.Context,
    filter model.SellerFilter,
) ([]model.SellerListItem, error) {

    query := `
        SELECT
            s.id,
            s.name,
            s.slug,
            s.ogrn_ogrnip,
            s.inn,
            COUNT(DISTINCT gi.sku) AS goods_amount,
            round(COALESCE(AVG(rv.score), 0), 2) AS average_review_score
        FROM parsing_data.seller s
        JOIN parsing_data.good_item gi
            ON gi.seller_id = s.id
        LEFT JOIN parsing_data.review rv
            ON rv.sku = gi.sku
    `

    var (
        where []string
        having []string
        args []any
        n = 1
    )

    if filter.BrandID != nil {
    		query += `
    		    JOIN parsing_data.brand_seller bs
    		        ON bs.seller_id = s.id
    		`
        where = append(where, fmt.Sprintf("bs.brand_id = $%d", n))
        args = append(args, *filter.BrandID)
        n++
    }

		if filter.CategoryID != nil {
		    query += `
		        JOIN parsing_data.good g
		            ON g.sku = gi.sku
		        JOIN parsing_data.category_relation cr
		            ON cr.child_id = g.cat_id
		    `
		
		    where = append(where, fmt.Sprintf("cr.parent_id = $%d", n))
		    args = append(args, *filter.CategoryID)
		    n++
		}

    if len(where) > 0 {
        query += "\nWHERE " + strings.Join(where, " AND ")
    }

    query += `
        GROUP BY
            s.id,
            s.name,
            s.slug,
            s.ogrn_ogrnip,
            s.inn
    `

    if filter.MinGoods != nil {
        having = append(
            having,
            fmt.Sprintf("COUNT(DISTINCT gi.sku) >= $%d", n),
        )
        args = append(args, *filter.MinGoods)
        n++
    }

    if filter.MinScore != nil {
        having = append(
            having,
            fmt.Sprintf("COALESCE(AVG(rv.score),0) >= $%d", n),
        )
        args = append(args, *filter.MinScore)
        n++
    }

    if len(having) > 0 {
        query += "\nHAVING " + strings.Join(having, " AND ")
    }

    query += `
        ORDER BY goods_amount DESC, s.name
    `

    rows, err := r.db.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    sellers := make([]model.SellerListItem, 0)

    for rows.Next() {
        var seller model.SellerListItem

        if err := rows.Scan(
            &seller.ID,
            &seller.Name,
            &seller.Slug,
            &seller.Ogrn,
            &seller.Inn,
            &seller.GoodsAmount,
            &seller.AverageReviewScore,
        ); err != nil {
            return nil, err
        }

        sellers = append(sellers, seller)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return sellers, nil
}

// ######################################################## SELLER BRAND
func (r *Repository) SellerBrands(
    ctx context.Context,
    sellerIDs []string,
) (map[string][]model.BrandSeller, error) {

    if len(sellerIDs) == 0 {
        return map[string][]model.BrandSeller{}, nil
    }

    rows, err := r.db.QueryContext(ctx, `
        SELECT
            bs.seller_id,
            b.id,
            b.slug,
            b.title
        FROM parsing_data.brand_seller bs
        JOIN parsing_data.brand b
            ON b.id = bs.brand_id
        WHERE bs.seller_id = ANY($1)
        ORDER BY b.title
    `, pq.Array(sellerIDs))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make(map[string][]model.BrandSeller)

    for rows.Next() {
        var (
            sellerID string
            brand model.BrandSeller
        )

        if err := rows.Scan(
            &sellerID,
            &brand.ID,
            &brand.Slug,
            &brand.Title,
        ); err != nil {
            return nil, err
        }

        result[sellerID] = append(result[sellerID], brand)
    }

    return result, rows.Err()
}

// ####################################################### SELLER GOODS
func (r *Repository) SellerGoods(
    ctx context.Context,
    sellerIDs []string,
) (map[string][]model.GoodSeller, error) {

    if len(sellerIDs) == 0 {
        return map[string][]model.GoodSeller{}, nil
    }

    rows, err := r.db.QueryContext(ctx, `
        SELECT
            seller_id,
            sku,
            slug,
            title
        FROM parsing_data.good_item
        WHERE seller_id = ANY($1)
        ORDER BY title
    `, pq.Array(sellerIDs))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make(map[string][]model.GoodSeller)

    for rows.Next() {
        var (
            sellerID string
            good model.GoodSeller
        )

        if err := rows.Scan(
            &sellerID,
            &good.Sku,
            &good.Slug,
            &good.Title,
        ); err != nil {
            return nil, err
        }

        result[sellerID] = append(result[sellerID], good)
    }

    return result, rows.Err()
}
