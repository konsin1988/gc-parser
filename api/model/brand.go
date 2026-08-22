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
}


type ResponseBrandById struct {
	Info										*BrandListItem		`json:"info"`
	Brands									[]SellersBrand		`json:"sellers"`
	Goods										[]GoodsBrand 			`json:"goods"`
}

type SellersBrand struct {
	ID						int	
	Slug					string
	Title					string
}

type GoodsBrand struct {
	Sku						string
	Title 				string
	Slug					string
}
