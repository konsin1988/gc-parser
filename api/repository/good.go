package repository

import (
    "context"
		_ "fmt"
		_ "strings"

		_ "github.com/lib/pq"

		"konsin1988/gc-api/model"
)


// ##################################################### SELLER LIST
func (r *Repository) GoodList(
    ctx context.Context,
    filter model.SellerFilter,
) ([]model.SellerListItem, error) {
	return nil, nil
}
