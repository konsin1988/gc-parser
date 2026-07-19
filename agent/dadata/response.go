package dadata

type FindByIDResponse struct {
    Suggestions []suggestion `json:"suggestions"`
}

type CompanyData struct {
    INN string `json:"inn"`
}

type suggestion struct {
    Value string  `json:"value"`
    Data  CompanyData `json:"data"`
}


type Company struct {
	Name string
	INN  string
}
