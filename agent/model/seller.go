package model


type Seller struct {
	ID				string	
	Name 			string	
	Slug 			string
	Ogrn			string
	Inn				string
}


type ParsedSeller struct {
    Name 			string
    OGRNIP 		string
    Slug			string
}
