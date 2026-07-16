package repository

import (
	"context"
)

func (r *Repository) InsertCategory(
	ctx context.Context,
	name string,
) (int, error) {

	var id int

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO parsing_data.category (name)
		VALUES ($1)
		ON CONFLICT (name)
		DO UPDATE SET
			name = EXCLUDED.name
		RETURNING id
	`, name).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
