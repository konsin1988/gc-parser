package service

import (
	"context"

	model "konsin1988/gc-api/model"
)

type SellerRepository interface {
	SellerList(ctx context.Context, filter model.SellerFilter) ([]model.SellerListItem, error)
	SellerBrands(ctx context.Context, sellerIDs []string) (map[string][]model.BrandSeller, error)
	SellerGoods(ctx context.Context, sellerIDs []string) (map[string][]model.GoodSeller, error)
}

type SellerService struct {
  repo SellerRepository 
}

func NewSellerService(repo SellerRepository) *SellerService {
  return &SellerService{repo: repo}
}

func (s *SellerService) SellerList(
    ctx context.Context,
    filter model.SellerFilter,
) ([]model.ResponseSellerListItem, error) {

    sellers, err := s.repo.SellerList(ctx, filter)
    if err != nil {
        return nil, err
    }

    sellerIDs := make([]string, 0, len(sellers))
    for _, seller := range sellers {
        sellerIDs = append(sellerIDs, seller.ID)
    }

    brandsBySeller, err := s.repo.SellerBrands(ctx, sellerIDs)
    if err != nil {
        return nil, err
    }

    goodsBySeller, err := s.repo.SellerGoods(ctx, sellerIDs)
    if err != nil {
        return nil, err
    }

    response := make([]model.ResponseSellerListItem, 0, len(sellers))

    for _, seller := range sellers {
        response = append(response, model.ResponseSellerListItem{
            ID:                 seller.ID,
            Name:               seller.Name,
            Slug:               seller.Slug,
            Ogrn:               seller.Ogrn,
            Inn:                seller.Inn,
            GoodsAmount:        seller.GoodsAmount,
            AverageReviewScore: seller.AverageReviewScore,
            Brands:             brandsBySeller[seller.ID],
            Goods:              goodsBySeller[seller.ID],
        })
    }

    return response, nil
}
