package model


type AllBrandsItem struct {
	ID 						string
	Title					string
	Slug					string
}


type BrandListItem struct {
	ID											string           `json:"id"` 
	Title										string					 `json:"name"`
	Slug										string					 `json:"slug"`
	GoodsAmount 						int     				 `json:"goodsAmount"`
	AverageReviewScore			float64 				 `json:"averageReviewScore"`
	SellerAmount						int							 `json:"sellerAmount"`
}


type ResponseBrandById struct {
	Info										*BrandListItem		`json:"info"`
	Brands									[]SellerBrand		`json:"sellers"`
	Goods										[]GoodBrand 			`json:"goods"`
}

type SellerBrand struct {
	ID						int	
	Slug					string
	Title					string
}

type GoodBrand struct {
	Sku						string
	Title 				string
	Slug					string
}
