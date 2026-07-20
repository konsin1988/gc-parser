package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) insertCategory(
	ctx context.Context,
	marketplace string,
	cat model.Category,
) (int, error) {

	var id int

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO parsing_data.category (marketplace, name, slug)
		VALUES ($1, $2, $3)
		ON CONFLICT (marketplace, slug)
		DO UPDATE SET
			name = EXCLUDED.name
		RETURNING id
	`, marketplace, cat.Name, cat.Slug).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) insertCategoryRelation(
    ctx context.Context,
    parentID,
    childID int,
) error {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO parsing_data.category_relation
		(parent_id, child_id)
		VALUES ($1,$2)
		ON CONFLICT (parent_id, child_id) 
		DO NOTHING;
	`, parentID, childID)

	return err 
}



func (r *Repository) InsertOzonCategories (
    ctx context.Context,
		breadcrumb []model.Category,
) (int, error) {

	var parentID int
	
	for _, cat := range breadcrumb {
	
	    id, err := r.insertCategory(
	        ctx,
	        "ozon",
	        cat,
	    )
	    if err != nil {
	        return 0, err
	    }
	
	    if parentID != 0 {
	        err = r.insertCategoryRelation(
	            ctx,
	            parentID,
	            id,
	        )
	        if err != nil {
	            return 0, err
	        }
	    }
	    parentID = id
	}
	return parentID, nil
}
