package main

import (
  "log"
	"context"
	"strconv"

  config "konsin1988/gc-agent/config"
	"konsin1988/gc-agent/dadata"
	"konsin1988/gc-agent/repository"
	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/etl"
	"konsin1988/gc-agent/cli"

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

	cmd, err := cli.ParseCLI()
	if err != nil {
		log.Fatal(err)
	}

	switch cmd.Command {
	case "search", "brand", "seller", "category":

		if cmd.Command == "search" {
			goodsJob := etl.NewSearchGoodsJob(
				services,
				cmd.Value,
				cmd.MaxPages,
			)

			if err := goodsJob.Run(ctx); err != nil {
				log.Fatal(err)
			}
		} else {
			itemId, err := strconv.Atoi(cmd.Value)
			if err != nil {
				log.Fatal(err)
			}

			goodsByItemPrefixJob := etl.NewGoodsByItemPrefixJob(
				services,
				cmd.Command,
				itemId,
				cmd.MaxPages,
			)

			if err := goodsByItemPrefixJob.Run(ctx); err != nil {
				log.Fatal(err)
			}
		}

	default:
	  log.Fatalf("unknown command: %s", cmd.Command)
	}
}

