package repository

import (
	"context"
	"database/sql"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertReviewImages (
	ctx context.Context,
	tx *sql.Tx,
	imgs []model.ReviewImage,
) (error) {

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO parsing_data.review_image (review_uuid, url)
		VALUES ($1, $2)
		ON CONFLICT (url)
		DO NOTHING; 
	`)
	if err != nil {
    return err
  }
  defer stmt.Close()


	for _, img := range imgs {
    _, err := stmt.ExecContext(
        ctx,
				img.ReviewUUID,
				img.ImgURL,
    )
  	if err != nil {
  		return err
    }
	}
	return nil 
}
