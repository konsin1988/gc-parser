package http

import (
  "net/http"
	"strings"
	"strconv"
	"fmt"

	"konsin1988/gc-api/model"
)


func parseFilter(r *http.Request) (model.Filter, error) {
    var filter model.Filter

    q := r.URL.Query()

		// brands
		for _, value := range q["brand"] {
        id, err := strconv.Atoi(value)
        if err != nil {
            return filter, fmt.Errorf("invalid brand: %s", value)
        }

        if id <= 0 {
            return filter, fmt.Errorf("invalid brand: %s", value)
        }

        filter.BrandIDs = append(filter.BrandIDs, id)
    }

		// categories
		for _, value := range q["category"] {
        id, err := strconv.Atoi(value)
        if err != nil {
            return filter, fmt.Errorf("invalid category: %s", value)
        }

        if id <= 0 {
            return filter, fmt.Errorf("invalid brand: %s", value)
        }

        filter.CategoryIDs = append(filter.CategoryIDs, id)
    }

    if v := q.Get("min_goods"); v != "" {
        n, err := strconv.Atoi(v)
        if err != nil {
            return filter, fmt.Errorf("invalid minGoods")
        }
        filter.MinGoods = &n
    }

		if v := q.Get("max_goods"); v != "" {
		    n, err := strconv.Atoi(v)
		    if err != nil {
		        return filter, fmt.Errorf("invalid max_goods")
		    }
        filter.MaxGoods = &n
		}

		// min > max good amount
		if filter.MinGoods != nil && 
			filter.MaxGoods != nil && 
			*filter.MinGoods > *filter.MaxGoods {
			return filter, fmt.Errorf("min_goods cannot be greater than max_goods")
		}


    if v := q.Get("max_score"); v != "" {
        score, err := strconv.ParseFloat(v, 64)
        if err != nil {
            return filter, fmt.Errorf("invalid maxScore")
        }
        filter.MaxScore = &score
    }

    if v := q.Get("min_score"); v != "" {
        score, err := strconv.ParseFloat(v, 64)
        if err != nil {
            return filter, fmt.Errorf("invalid minScore")
        }
        filter.MinScore = &score
    }

		if v := q.Get("sort"); v != "" {
        sort, err := parseSort(v)
        if err != nil {
            return filter, err
        }

        filter.Sort = sort
    }

    return filter, nil
}


// ---------------------------------------------- parseSellerSort
func parseSort(value string) ([]model.Sort, error) {
    parts := strings.Split(value, ",")

    result := make([]model.Sort, 0, len(parts))

    for _, part := range parts {
        pieces := strings.Split(part, ":")

        if len(pieces) != 2 {
            return nil, fmt.Errorf(
                "invalid sort %q, expected field:order",
                part,
            )
        }

        field := pieces[0]
        order := pieces[1]

        switch field {
        case "name", "goods", "score":
        default:
            return nil, fmt.Errorf("invalid sort field: %s", field)
        }

        switch order {
        case "asc", "desc":
        default:
            return nil, fmt.Errorf("invalid sort order: %s", order)
        }

        result = append(result, model.Sort{
            Field: field,
            Order: order,
        })
    }

    return result, nil
}
