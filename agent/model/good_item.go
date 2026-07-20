package model

type GoodItem struct {
	Sku 							string	
	Slug							string
	Title							string	
	Price							*int
	CardPrice					*int
	OriginalPrice			*int
	Availability			bool	
	SellerId					*string
	BrandId						*string
	ReviewLink				string
}


type ParsedGoodItem struct {
	Sku 							string	
	Title							string	
	Price							*int
	CardPrice					*int
	OriginalPrice			*int
	Availability			bool	
	ReviewLink				string

	Categories				[]Category
	Brand							*Brand
	Images						[]Image
	Seller						*ParsedSeller
}
