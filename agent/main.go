package main

import (
  _ "net/http"
  "log"
	"context"

  config "konsin1988/gc-agent/config"
	"konsin1988/gc-agent/dadata"
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

	dadataClient := dadata.New(config.App)
	repo := repository.New(db)
	ozonClient, err := ozon.New(config.App)
	if err != nil {
		log.Fatal(err)
	}

	services := &etl.Services{
		Repo:   repo,
		Ozon:   ozonClient,
		Dadata: dadataClient,
	}

	queryID, err := repo.InsertQuery(
		ctx,
		config.App.Ozon.SearchQuery,
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// goodsJob
	goodsJob := etl.NewSearchGoodsJob(
		services,
		config.App.Ozon.SearchQuery,
		queryID,
		5,
	)

	if err := goodsJob.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

