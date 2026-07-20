package model

type Image struct {
	Sku 			string	
	ImgURL		string	
	IsCover		bool
}

type ReviewImage struct {
	ReviewUUID 			string
	ImgURL					string
}
