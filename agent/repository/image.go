package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertImage (
	ctx context.Context,
	imgs []model.Image,
) (error) {

	tx, err := r.db.BeginTx(ctx, nil)
  if err != nil {
      return err
  }
  defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO parsing_data.image (sku, img_url, is_cover)
		VALUES ($1, $2, $3)
		ON CONFLICT (img_url)
		DO UPDATE SET
			img_url = EXCLUDED.img_url
	`)
	if err != nil {
    return err
  }
  defer stmt.Close()


	for _, img := range imgs {
    _, err := stmt.ExecContext(
        ctx,
				img.Sku,
				img.ImgURL,
				img.IsCover,
    )
  	if err != nil {
  		return err
    }
	}
	return tx.Commit() 
}
