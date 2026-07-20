package etl

import (
	"context"

	"konsin1988/gc-agent/marketplace/ozon"
	"konsin1988/gc-agent/parser"
)

type ReviewJob struct {
   	Services 

    ReviewURL			string 
		MaxPages			int
}



func (j *ReviewJob) Fetch(ctx context.Context) (any, error) {
    return j.Ozon.DataByURL(ctx, j.ReviewURL)
}

func (j *ReviewJob) Parse(data any) (any, error) {
    page := data.(*ozon.PageResponse)

    return parser.ParseReview(page)
}

func (j *ReviewJob) Save(ctx context.Context, data any) error {

   parsed := data.(*parser.ReviewPage)
	 err := j.Repo.InsertReviews (
       ctx,
			 parsed.Reviews,
   )
	 return err
}

func NewReviewJob(
	services *Services,
	reviewURL string,
	maxPages	int,
) *ReviewJob {
	return &ReviewJob{
		Services: *services,
		ReviewURL: reviewURL,
		MaxPages: maxPages,
	}
}


func (j *ReviewJob) Run(ctx context.Context) error {

    for i := 0; i < j.MaxPages; i++ {

        raw, err := j.Fetch(ctx)
        if err != nil {
            return err
        }

        parsed, err := j.Parse(raw)
        if err != nil {
            return err
        }

        page := parsed.(*parser.ReviewPage)

        if err := j.Save(ctx, page); err != nil {
            return err
        }

        if page.NextPage == "" {
            break
        }

        j.ReviewURL = page.NextPage
    }

    return nil
}
