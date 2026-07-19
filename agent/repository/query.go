package repository

import (
	"context"
)

func (r *Repository) InsertQuery(
	ctx context.Context,
	query string,
) (int, error) {

	var id int

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO parsing_data.query (query)
		VALUES ($1)
		ON CONFLICT (query)
		DO UPDATE SET
			query = EXCLUDED.query
		RETURNING id
	`, query).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
