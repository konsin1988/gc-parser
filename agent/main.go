package main

import (
  _ "net/http"
  "log"
	"context"

  config "konsin1988/gc-agent/config"
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

	repo := repository.New(db)

	ctx := context.Background()
	
	catID, err := repo.InsertCategory(
		ctx,
		config.App.Ozon.SearchCategory,
	)
	if err != nil {
		log.Fatal(err)
	}
	
	queryID, err := repo.InsertQuery(
		ctx,
		config.App.Ozon.SearchQuery,
	)
	if err != nil {
		log.Fatal(err)
	}
	
	ozonClient, err := ozon.New(config.App)
	if err != nil {
		log.Fatal(err)
	}

	goodsJob := etl.NewSearchGoodsJob(
		ozonClient,
		repo,
		config.App.Ozon.SearchQuery,
		catID,
		queryID,
		5,
	)

	if err := goodsJob.Run(ctx); err != nil {
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

