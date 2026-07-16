package config

import (
    "os"
)

type OzonConfig struct {
	APIBaseURL 				string
	Cookie						string
	SearchURL					string
	SearchQuery				string
	SearchCategory		string
	SingleItemURL			string
}

type DadataConfig struct {
	DadataApiToken		string
	DadataApiURL			string
}

type HeadersConfig struct {
	UserAgent					string
	AppName						string
  AppVersion				string
  ManifestVersion		string
  ParentRequestID		string
  PageViewID				string
}

type Config struct {
	Headers 		HeadersConfig
	Ozon 				OzonConfig
	Dadata			DadataConfig
}

var App Config

func Load() {
	App = Config{
		Headers: HeadersConfig{
			UserAgent: 				os.Getenv("HEADERS_USER_AGENT"),
			AppName:					os.Getenv("HEADERS_X_O3_APP_NAME"),
			AppVersion:				os.Getenv("HEADERS_X_O3_APP_VERSION"),
			ManifestVersion: 	os.Getenv("HEADERS_X_O3_MANIFEST_VERSION"),
			ParentRequestID: 	os.Getenv("HEADERS_X_O3_PARENT_REQUESTID"),
			PageViewID:				os.Getenv("HEADERS_X_PAGE_VIEW_ID"),
		},
		Ozon: OzonConfig{
			APIBaseURL: 			os.Getenv("OZON_API_BASE_URL"),
			Cookie:     			os.Getenv("OZON_COOKIE"),
			SearchURL: 				os.Getenv("OZON_SEARCH_URL"),
			SearchQuery: 			os.Getenv("OZON_SEARCH_QUERY"),
			SearchCategory:		os.Getenv("OZON_SEARCH_CATEGORY"),
			SingleItemURL:		os.Getenv("OZON_SINGLE_ITEM_URL"),
		},
		Dadata: DadataConfig{
			DadataApiToken:		os.Getenv("DADATA_API_TOKEN"),			
			DadataApiURL:			os.Getenv("DADATA_API_URL"),
		},
	}
}
