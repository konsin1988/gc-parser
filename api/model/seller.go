package model

type SellerFilter struct {
  BrandIDs			[]int 
  CategoryIDs   []int
  MinGoods     	*int
	MaxGoods 			*int
  MinScore     	*float64
	MaxScore			*float64	

	Sort				[]SellerSort
}

type SellerSort struct {
	Field			string
	Order			string
}

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

