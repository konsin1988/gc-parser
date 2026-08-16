package model

import (
	"time"
)

type GoodFilter struct {
	Brands				[]int
	Sellers				[]int
	Price					[]int
	Availability bool
  MinScore     *float64

	Sort				[]GoodSort
}

type GoodSort struct {
	Field			string
	Order			string
}


type ResponseGoodById  struct {
	Sku 										string            `json:"sku"`
	Title										string						`json:"title"`	
  Slug										string						`json:"slug"`
	Price 									int     					`json:"price"`
	CardPrice 							int     					`json:"card_price"`
	OriginalPrice 					int     					`json:"original_price"`
	Availability						bool							`json:"availability"`	
	Sellers									[]SellerGood			`json:"sellers"`
	Brand										BrandGood					`json:"brand"`
	Reviews									[]ReviewGood			`json:"reviews"`
}


type GoodListItem struct {
	Sku 										string            `json:"sku"`
	Title										string						`json:"title"`	
  Slug										string						`json:"slug"`
	Price 									int     					`json:"price"`
	CardPrice 							int     					`json:"card_price"`
	OriginalPrice 					int     					`json:"original_price"`
	Availability						bool							`json:"availability"`	
}

type SellerGood struct {
	ID					string
	Name				string
	Slug				string
}

type BrandGood struct {
	ID					int
	Name				string
	Slug				string
}

type ReviewGood	struct {
	UUID				string
	CreatedAt		time.Time
	Score				int
	Comment			string
}

type CategoryGood struct {
	Name 				string
	Slug				string
	ID					string
	ParentID		string
}

