package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertReviews(
	ctx context.Context,
	reviews []model.Review,
) (error) {

	tx, err := r.db.BeginTx(ctx, nil)
  if err != nil {
      return err
  }
  defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO parsing_data.review(uuid, created_at, sku, author_guid, comment, positive, negative)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (uuid)
		DO UPDATE SET
			comment = EXCLUDED.comment,
			positive = EXCLUDED.positive,
			negative = EXCLUDED.negative
	`)
	if err != nil {
    return err
  }
  defer stmt.Close()


	for _, review := range reviews {
    _, err := stmt.ExecContext(
        ctx,
				review.UUID,
				review.CreatedAt,
				review.Sku,
				review.AuthorGuid,
				review.Comment,
				review.Positive,
				review.Negative,
    )
  	if err != nil {
  		return err
    }

		if len(review.ReviewImages) > 0 {
			err = r.InsertReviewImages(ctx, tx, review.ReviewImages)
			if err != nil {
				return err
			}
		}
	}
	return tx.Commit() 
}
