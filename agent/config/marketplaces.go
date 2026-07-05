package config

import (
    "os"
)

type OzonConfig struct {
	APIBaseURL 				string
	Cookie						string
	SearchURL					string
	SearchQuery				string
	SingleItemURL			string
}

type DadataConfig struct {
	DadataApiToken		string
	DadataApiURL			string
}

type Config struct {
	Ozon 		OzonConfig
	Dadata	DadataConfig
}

var App Config

func Load() {
	App = Config{
		Ozon: OzonConfig{
			APIBaseURL: 			os.Getenv("OZON_API_BASE_URL"),
			Cookie:     			os.Getenv("OZON_COOKIE"),
			SearchURL: 				os.Getenv("OZON_SEARCH_URL"),
			SearchQuery: 			os.Getenv("OZON_SEARCH_QUERY"),
			SingleItemURL:		os.Getenv("OZON_SINGLE_ITEM_URL"),
		},
		Dadata: DadataConfig{
			DadataApiToken:		os.Getenv("DADATA_API_TOKEN"),			
			DadataApiURL:			os.Getenv("DADATA_API_URL"),
		},
	}
}
