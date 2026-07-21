package service

import (
	"context"

	"konsin1988/gc-agent/model"
)


func (s *GoodItemService) ProcessGoodItem(
    ctx context.Context,
    parsed *model.ParsedGoodItem,
		goodURL string,
		queryID	int,
) error {
	var brandID int
	var err error

	// --------------------------------------------------  Brand
	if parsed.Brand != nil{	
		brandID, err = s.Repo.InsertBrand(ctx, *parsed.Brand) 
		if err != nil {
			return err
		}
	}	

	// ---------------------------------------------------  InsertSeller 
	seller := model.Seller{}	
	if parsed.Seller.ID == "0"{
		seller = model.Seller{
			ID: "0",
			Name: "ООО \"Интернет Решения\"",
			Slug: "ozon",
			Ogrn: "1027739244741",
			Inn: "7704217370",
		}
	} else { 
		company, err := s.Dadata.FindByOGRNIP(
			ctx, 
			parsed.Seller.OGRNIP,
		)
		if err != nil {
			return err
		}
		seller = model.Seller{
			ID: parsed.Seller.ID,	
			Name: parsed.Seller.Name, 
			Slug: parsed.Seller.Slug,
			Ogrn: parsed.Seller.OGRNIP,
			Inn: company.INN,	
		}
	}
	err = s.Repo.InsertSeller(ctx, seller)
	if err != nil{
		return err
	}

	// --------------------------------------------------------------  InsertOzonCategories
	catID, err := s.Repo.InsertOzonCategories(
		ctx, 
		parsed.Categories,
	)
	if err != nil {
		return err
	}

	// ------------------------------------------  InsertGoodItem
	goodSlug := GoodSlug(goodURL) 

	goodItem := model.GoodItem{
		Sku 					: parsed.Sku, 
		Slug					: goodSlug,
		Title					: parsed.Title, 
		Price					: parsed.Price,
		CardPrice			: parsed.CardPrice,
		OriginalPrice	: parsed.OriginalPrice,
		Availability	: parsed.Availability,
		SellerId			: parsed.Seller.ID,
		BrandId				:	brandID, 
		ReviewLink		: parsed.ReviewLink,
	}
	err = s.Repo.InsertGoodItem(ctx, &goodItem)
	if err != nil {
		return err
	}

	// -------------------------------------------  InsertImage
	err = s.Repo.InsertImage(ctx, parsed.Images)
	if err != nil{
		return err
	}

	// --------------------------------------------  InsertGood
	good := model.Good{
		Sku: parsed.Sku,
		Link: goodURL,
	}	
	err = s.Repo.InsertGood(ctx, catID, queryID, good)
	if err != nil {
		return err
	}


	return nil
}
