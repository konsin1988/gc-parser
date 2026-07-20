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
	BrandId						*int
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

// we need to: 
// 1. Save Brand and receive it's id.
// 2. Save images 
// 3. Save seller 
// 4. Save Categories and receive cat_id 
// 5. Get and save seller's brands (SellerJob with SellerID) 
// 6. Save GoodItem 
// 7. Save Good in goods. 
// 8. Get and save reviews (ReviewJob with ReviewLink)
