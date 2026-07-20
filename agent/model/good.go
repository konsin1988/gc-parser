package model

type Good struct {
	Sku 			string	
	Link 			string	
}


type GoodsPage struct {
	Goods 			[]Good
	NextPage		string
}
