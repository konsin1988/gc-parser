package etl

//import (
//	"context"
//
//	"konsin1988/gc-agent/repository"
//	"konsin1988/gc-agent/marketplace/ozon"
//	"konsin1988/gc-agent/parser"
//	"konsin1988/gc-agent/model"
//)
//
//
//type GoodItemJob struct {
//    BaseJob
//
//    QueryID int
//    Link    string
//}
//
//type ParsedGood struct {
//    Good       model.GoodItem
//    Breadcrumb []model.Category
//    Brand      *model.Brand
//    Images     []model.Image
//}
//
//func (j *GoodItemJob) Fetch(ctx context.Context) (any, error) {
//    return j.Client.GoodByLink(ctx, j.Link)
//}
