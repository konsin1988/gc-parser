package service

import (
	"context"

	model "konsin1988/gc-api/model"
)

type SellerRepository interface {
	AllSellers(ctx context.Context) ([]model.AllSellersItem, error)
	SellerList(ctx context.Context, filter model.SellerFilter) ([]model.SellerListItem, error)
	SellerById(ctx context.Context, SellerID string) (*model.SellerListItem, error)
	SellerBrands(ctx context.Context, sellerIDs []string) (map[string][]model.BrandSeller, error)
	SellerGoods(ctx context.Context, sellerIDs []string) (map[string][]model.GoodSeller, error)
}

type SellerService struct {
  repo SellerRepository 
}

func NewSellerService(repo SellerRepository) *SellerService {
  return &SellerService{repo: repo}
}

func (s *SellerService) AllSellers(
    ctx context.Context,
) ([]model.AllSellersItem, error) {

    sellers, err := s.repo.AllSellers(ctx)
    if err != nil {
        return nil, err
    }
		return sellers, nil
}


func (s *SellerService) SellerList(
    ctx context.Context,
    filter model.SellerFilter,
) ([]model.SellerListItem, error) {

    sellers, err := s.repo.SellerList(ctx, filter)
    if err != nil {
        return nil, err
    }
		return sellers, nil
}

func (s *SellerService) SellerById(
    ctx context.Context,
		sellerID string,
) (*model.ResponseSellerById, error) {

    seller, err := s.repo.SellerById(ctx, sellerID)
    if err != nil {
        return nil, err
    }

		sellerIdList := []string{seller.ID}

    brandsBySeller, err := s.repo.SellerBrands(ctx, sellerIdList)
    if err != nil {
        return nil, err
    }

    goodsBySeller, err := s.repo.SellerGoods(ctx, sellerIdList)
    if err != nil {
        return nil, err
    }

		response := &model.ResponseSellerById{
		        ID:                 seller.ID,
		        Name:               seller.Name,
		        Slug:               seller.Slug,
		        Ogrn:               seller.Ogrn,
		        Inn:                seller.Inn,
		        GoodsAmount:        seller.GoodsAmount,
		        AverageReviewScore: seller.AverageReviewScore,
		        Brands:             brandsBySeller[seller.ID],
		        Goods:              goodsBySeller[seller.ID],
		    }
		
		return response, nil
}
