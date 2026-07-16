package etl

type SellerJob struct {
    BaseJob

    SellerID int64
}

//func (j *SellerJob) Fetch(ctx context.Context) (any, error) {
//    return j.Client.Seller(ctx, j.SearchText)
//}
//
//func (j *SellerJob) Parse(data any) (any, error) {
//    page := data.(*ozon.PageResponse)
//
//    return parser.ParseGoods(page, j.GridNum)
//}
//
//func (j *SellerJob) Save(ctx context.Context, data any) error {
//    goods := data.([]model.Good)
//
//    return j.Repo.SaveGoods(ctx, goods)
//}
