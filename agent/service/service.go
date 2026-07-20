package service


import (

	"konsin1988/gc-agent/repository"
	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/dadata"
)


type GoodItemService struct {
    Repo 			*repository.Repository
    Ozon       *ozon.Client
		Dadata		*dadata.Client
}

func NewGoodItemService(
	repo *repository.Repository,
	ozon *ozon.Client,
	dadata *dadata.Client,
) *GoodItemService {
	return &GoodItemService{
		Repo:   repo,
		Ozon:   ozon,
		Dadata: dadata,
	}
}
