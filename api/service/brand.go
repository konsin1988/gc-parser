package service

import (
	"context"

	model "konsin1988/gc-api/model"
)

type BrandRepository interface {
	AllBrands(ctx context.Context) ([]model.AllBrandsItem, error)
	//SellerList(ctx context.Context, filter model.Filter) ([]model.SellerListItem, error)
	//SellerById(ctx context.Context, SellerID string) (*model.SellerListItem, error)
	//SellerBrands(ctx context.Context, sellerIDs []string) (map[string][]model.BrandSeller, error)
	//SellerGoods(ctx context.Context, sellerIDs []string) (map[string][]model.GoodSeller, error)
}

type BrandService struct {
  repo BrandRepository 
}

func NewBrandService(repo BrandRepository) *BrandService {
  return &BrandService{repo: repo}
}

func (s *BrandService) AllBrands(
    ctx context.Context,
) ([]model.AllBrandsItem, error) {

    brands, err := s.repo.AllBrands(ctx)
    if err != nil {
        return nil, err
    }
		return brands, nil
}

