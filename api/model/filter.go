package model


type Filter struct {
  BrandIDs			[]int 
  CategoryIDs   []int
  MinGoods     	*int
	MaxGoods 			*int
  MinScore     	*float64
	MaxScore			*float64	

	Sort				[]Sort
}

type Sort struct {
	Field			string
	Order			string
}
