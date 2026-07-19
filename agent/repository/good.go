package repository

import (
	"context"

	"konsin1988/gc-agent/model"
)

func (r *Repository) InsertGood (
	ctx context.Context,
	catID int,
	queryID int,
	good model.Good,
) (error) {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO parsing_data.good (sku, cat_id, query_id, glink)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (sku, cat_id, query_id)
		DO UPDATE SET
			glink = EXCLUDED.glink
	`, good.Sku, catID, queryID, good.Link)

	return err 
}



//import (
//	"context"
//
//	"konsin1988/gc-agent/model"
//)
//
//func (r *Repository) SaveGoods(
//	ctx context.Context, 
//	catID int,
//	queryID int,
//	goods []model.Good,
//) error {
//
//	tx, err := r.db.BeginTx(ctx, nil)
//  if err != nil {
//      return err
//  }
//  defer tx.Rollback()
//
//	stmt, err := tx.PrepareContext(ctx, `
//		INSERT INTO parsing_data.good (
//        sku,
//        cat_id,
//				query_id,
//        glink
//    )
//    VALUES ($1, $2, $3, $4)
//    ON CONFLICT (good_id, cat_id, query_id)
//    DO UPDATE SET
//      glink = EXCLUDED.glink
//	`)
//	if err != nil {
//    return err
//  }
//  defer stmt.Close()
//
//	for _, good := range goods {
//    _, err := stmt.ExecContext(
//        ctx,
//        good.Sku,
//        catID,
//				queryID,
//        good.Link,
//    )
//  	if err != nil {
//  		return err
//    }
//	}
//	return tx.Commit() 
//}

