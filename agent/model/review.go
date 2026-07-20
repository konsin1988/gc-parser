package model

import "time"

type Review struct {
	UUID 						string
	CreatedAt				time.Time	
	Sku							string
	AuthorGuid			string
	Score						int
	Comment					string
	Positive				string
	Negative				string
	ReviewImages		[]ReviewImage	
}
