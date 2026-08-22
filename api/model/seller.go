package model


type AllSellersItem struct {
	ID 						string
	Name					string
	Slug					string
}

type SellerListItem struct {
	ID											string           `json:"id"` 
	Name										string					 `json:"name"`
	Slug										string					 `json:"slug"`
	Ogrn										string					 `json:"ogrn"`
	Inn											string					 `json:"inn"`
	GoodsAmount 						int     				 `json:"goodsAmount"`
	AverageReviewScore			float64 				 `json:"averageReviewScore"`
}

type ResponseSellerById struct {
	Info										*SellerListItem		`json:"info"`
	Brands									[]BrandSeller			`json:"brands"`
	Goods										[]GoodSeller 			`json:"goods"`
}

type BrandSeller struct {
	ID						int	
	Slug					string
	Title					string
}

type GoodSeller struct {
	Sku						string
	Title 				string
	Slug					string
}

