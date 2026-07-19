package main

import (
  _ "net/http"
  "log"
	"context"

  config "konsin1988/gc-agent/config"
	_ "konsin1988/gc-agent/dadata"
	"konsin1988/gc-agent/repository"
	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/etl"
	_ "github.com/bogdanfinn/fhttp"
  _ "github.com/bogdanfinn/tls-client"
  _ "github.com/bogdanfinn/tls-client/profiles"
)

func main() {
	config.Load()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

//	dadataClient := dadata.New(config.App)

	repo := repository.New(db)

	
	//queryID, err := repo.InsertQuery(
	//	ctx,
	//	config.App.Ozon.SearchQuery,
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	
	ozonClient, err := ozon.New(config.App)
	if err != nil {
		log.Fatal(err)
	}

	// goodsJob
	//goodsJob := etl.NewSearchGoodsJob(
	//	ozonClient,
	//	dadataClient,
	//	repo,
	//	config.App.Ozon.SearchQuery,
	//	queryID,
	//	5,
	//)

	//if err := goodsJob.Run(ctx); err != nil {
	//	log.Fatal(err)
	//}
	
	// sellerJob
	sellerJob := etl.NewSellerJob(
		ozonClient,
		repo,
		"43306",
	)

	if err := sellerJob.Run(ctx); err != nil {
		log.Fatal(err)
	}


	// reviewJob 
	reviewJob := etl.NewReviewJob(
		ozonClient,
		repo,
		"/product/klaviatura-provodnaya-logitech-k280e-black-usb-103-klavishi-vodostoykaya-920-005215-260147596/reviews/",
		5,
	)

	if err := reviewJob.Run(ctx); err != nil {
		log.Fatal(err)
	}


	//job := &etl.GetSellerJob{
	//    Client: ozonClient,
	//    Repo: repo,
	//    SellerID: 12345,
	//}
	//
	//err := etl.Run(ctx, job)

}

