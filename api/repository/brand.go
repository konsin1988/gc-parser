package repository

import (
    "context"
		"fmt"
		"strings"
		"errors"
		"database/sql"

		"github.com/lib/pq"

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


// ###################################################### BRAND BY ID
func (r *Repository) BrandById(
    ctx context.Context,
		brandID string,
) (*model.BrandListItem, error) {

    query := `
			with seller_count as (
			select 
				b.id,
            	b.title,
            	b.slug,
            	count(distinct bs.seller_id ) as seller_amount
            from parsing_data.brand b 
            right join parsing_data.brand_seller bs 
            on b.id = bs.brand_id 
            group by b.id, b.title, b.slug
			),
			review_count as (
				select 
					gi.sku,
					gi.brand_id,
					r.score 
				from parsing_data.good_item gi 
				left join parsing_data.review r 
				on gi.sku = r.sku 
			),
			good_count as (
				select 
					sc.id,
					sc.title,
					sc.slug,
					sc.seller_amount,
					coalesce(count(gi.sku), 0) as good_amount
				from seller_count sc
				right join parsing_data.good_item gi 
				on sc.id = gi.brand_id 
				where sc.id is not null
				group by sc.id, sc.title, sc.slug, sc.seller_amount 
			),
			result as (
				select 
					gc.*,
					round(COALESCE(AVG(rc.score), 0), 2) AS average_review_score
				from good_count gc
				right join review_count rc 
				on gc.id = rc.brand_id 
				where gc.id is not null
				group by gc.id, gc.title, gc.slug, seller_amount, gc.good_amount
			)
			select  
				r.id,
				r.title,
				r.slug,
				r.good_amount,
				r.average_review_score,
				r.seller_amount
			from result r
			where r.id = $1
    `
		var brand model.BrandListItem

    err := r.db.QueryRowContext(ctx, query, brandID).Scan(
        &brand.ID,
        &brand.Title,
        &brand.Slug,
        &brand.GoodsAmount,
        &brand.AverageReviewScore,
				&brand.SellerAmount,
    )

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }

    return &brand, nil
}


// ######################################################## BRAND SELLERS
func (r *Repository) BrandSellers(
    ctx context.Context,
    brandIDs []string,
) (map[string][]model.SellerBrand, error) {

    if len(brandIDs) == 0 {
        return map[string][]model.SellerBrand{}, nil
    }

    rows, err := r.db.QueryContext(ctx, `
        SELECT
            bs.brand_id,
            s.id,
            s.slug,
            s.name
        FROM parsing_data.brand_seller bs
        JOIN parsing_data.seller s 
            ON s.id = bs.seller_id
        WHERE bs.brand_id = ANY($1)
        ORDER BY s.name
    `, pq.Array(brandIDs))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make(map[string][]model.SellerBrand)

    for rows.Next() {
        var (
            brandID string
            seller model.SellerBrand
        )

        if err := rows.Scan(
            &brandID,
            &seller.ID,
            &seller.Slug,
            &seller.Title,
        ); err != nil {
            return nil, err
        }

        result[brandID] = append(result[brandID], seller)
    }

    return result, rows.Err()
}

// ####################################################### SELLER GOODS
func (r *Repository) BrandGoods(
    ctx context.Context,
    brandIDs []string,
) (map[string][]model.GoodBrand, error) {

    if len(brandIDs) == 0 {
        return map[string][]model.GoodBrand{}, nil
    }

    rows, err := r.db.QueryContext(ctx, `
        SELECT
            brand_id,
            sku,
            slug,
            title
        FROM parsing_data.good_item
        WHERE brand_id = ANY($1)
        ORDER BY title
    `, pq.Array(brandIDs))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make(map[string][]model.GoodBrand)

    for rows.Next() {
        var (
            brandID string
            good model.GoodBrand
        )

        if err := rows.Scan(
            &brandID,
            &good.Sku,
            &good.Slug,
            &good.Title,
        ); err != nil {
            return nil, err
        }

        result[brandID] = append(result[brandID], good)
    }

    return result, rows.Err()
}


// ##################################################### BRAND LIST
func (r *Repository) BrandList(
    ctx context.Context,
    filter model.Filter,
) ([]model.BrandListItem, error) {

    query := `
			select 
				b.id,
				b.title,
				b.slug,
				count(distinct gi.sku) as goods_amount,
				round(COALESCE(AVG(r.score), 0), 2) AS average_review_score,
				count(distinct bs.seller_id) as seller_amount
			from parsing_data.brand b 
			join parsing_data.brand_seller bs 
			on b.id = bs.brand_id 
			left join parsing_data.good_item gi 
			on gi.brand_id = b.id
			left join parsing_data.review r 
			on gi.sku = r.sku 
    `

    var (
        where []string
        having []string
        args []any
        n = 1
    )


		if len(filter.SellerIDs) > 0 {
		
		    where = append(
		        where,
		        fmt.Sprintf("bs.seller_id = ANY($%d)", n),
		    )
		
		    args = append(args, pq.Array(filter.SellerIDs))
		    n++
		}

		if len(filter.CategoryIDs) > 0 {
		    query += `
		        left JOIN parsing_data.good g
		            ON g.sku = gi.sku
		        left JOIN parsing_data.category_relation cr
		            ON cr.child_id = g.cat_id
		    `
		
		    where = append(
					where, 
					fmt.Sprintf("cr.parent_id = ANY($%d)", n),
				)
		    args = append(args, pq.Array(filter.CategoryIDs))
		    n++
		}

    if len(where) > 0 {
        query += "\nWHERE " + strings.Join(where, " AND ")
    }

    query += `
			group by b.id, b.title, b.slug
    `

    if filter.MinGoods != nil {
        having = append(
            having,
            fmt.Sprintf("COUNT(DISTINCT gi.sku) >= $%d", n),
        )
        args = append(args, *filter.MinGoods)
        n++
    }

		if filter.MaxGoods != nil {
		    having = append(
		        having,
		        fmt.Sprintf("COUNT(DISTINCT gi.sku) <= $%d", n),
		    )
		    args = append(args, *filter.MaxGoods)
		    n++
		}

    if filter.MaxScore != nil {
        having = append(
            having,
            fmt.Sprintf("COALESCE(AVG(r.score),5) <= $%d", n),
        )
        args = append(args, *filter.MaxScore)
        n++
    }

    if filter.MinScore != nil {
        having = append(
            having,
            fmt.Sprintf("COALESCE(AVG(r.score),0) >= $%d", n),
        )
        args = append(args, *filter.MinScore)
        n++
    }


    if len(having) > 0 {
        query += "\nHAVING " + strings.Join(having, " AND ")
    }


		sortColumns := map[string]string{
		    "name":  "b.title",
		    "goods": "goods_amount",
		    "score": "average_review_score",
		}

		sortParts := make([]string, 0, len(filter.Sort))
		
		for _, sort := range filter.Sort {
		    column, ok := sortColumns[sort.Field]
		    if !ok {
		        return nil, fmt.Errorf("invalid sort field: %s", sort.Field)
		    }
		
		    direction := "ASC"
		
		    if sort.Order == "desc" {
		        direction = "DESC"
		    }
		
		    sortParts = append(
		        sortParts,
		        fmt.Sprintf("%s %s", column, direction),
		    )
		}

		if len(sortParts) > 0 {
		    query += "\nORDER BY " + strings.Join(sortParts, ", ")
		} else {
		    query += "\nORDER BY goods_amount DESC, b.id"
		}


    rows, err := r.db.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    brands := make([]model.BrandListItem, 0)

    for rows.Next() {
        var brand model.BrandListItem

        if err := rows.Scan(
            &brand.ID,
            &brand.Title,
            &brand.Slug,
            &brand.GoodsAmount,
            &brand.AverageReviewScore,
						&brand.SellerAmount,
        ); err != nil {
            return nil, err
        }

        brands = append(brands, brand)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return brands, nil
}
