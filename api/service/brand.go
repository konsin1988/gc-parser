package service

import (
	"context"

	model "konsin1988/gc-api/model"
)

type BrandRepository interface {
	AllBrands(ctx context.Context) ([]model.AllBrandsItem, error)
	BrandList(ctx context.Context, filter model.Filter) ([]model.BrandListItem, error)
	BrandById(ctx context.Context, BrandID string) (*model.BrandListItem, error)
	BrandSellers(ctx context.Context, brandIDs []string) (map[string][]model.SellerBrand, error)
	BrandGoods(ctx context.Context, brandIDs []string) (map[string][]model.GoodBrand, error)
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

func (s *BrandService) BrandById(
    ctx context.Context,
		brandID string,
) (*model.ResponseBrandById, error) {

    brand, err := s.repo.BrandById(ctx, brandID)
    if err != nil {
        return nil, err
    }

		brandIdList := []string{brand.ID}

    sellersByBrand, err := s.repo.BrandSellers(ctx, brandIdList)
    if err != nil {
        return nil, err
    }

    goodsByBrand, err := s.repo.BrandGoods(ctx, brandIdList)
    if err != nil {
        return nil, err
    }

		response := &model.ResponseBrandById{
						Info:								brand,
		        Brands:             sellersByBrand[brand.ID],
		        Goods:              goodsByBrand[brand.ID],
		    }

		return response, nil
}


func (s *BrandService) BrandList(
    ctx context.Context,
    filter model.Filter,
) ([]model.BrandListItem, error) {

    brands, err := s.repo.BrandList(ctx, filter)
    if err != nil {
        return nil, err
    }
		return brands, nil
}
