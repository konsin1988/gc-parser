package model

type SellerFilter struct {
  BrandID      *int
  CategoryID   *int
  MinGoods     *int
  MinScore     *float64
}


type SellerListItem struct {
	ID											string           
	Name										string					 
  Slug										string					 
  Ogrn										string					 
  Inn											string					 
	GoodsAmount 						int     				 
	AverageReviewScore			float64 				 
}

type ResponseSellerListItem struct {
	ID											string            `json:"id"`
	Name										string						`json:"name"`	
  Slug										string						`json:"slug"`
  Ogrn										string						`json:"ogrn"`
  Inn											string						`json:"inn"`
	GoodsAmount 						int     					`json:"goodsAmount"`
	Brands									[]BrandSeller			`json:"brands"`
	Goods										[]GoodSeller 			`json:"goods"`
	AverageReviewScore			float64 					`json:"averageReviewScore"`
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
