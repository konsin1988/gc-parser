package http

import (
  "encoding/json"
  "net/http"
	"context"
	"strings"
	"strconv"
	"errors"
	"fmt"
	"log"

	"konsin1988/gc-api/model"
)

type SellerService interface {
		AllSellers(ctx context.Context) ([]model.AllSellersItem, error)
    SellerList(ctx context.Context, filter model.SellerFilter) ([]model.SellerListItem, error)
		SellerById(ctx context.Context, SellerId string) (*model.ResponseSellerById, error)
}

type SellerHandler struct {
    service SellerService
}

func NewSellerHandler(service SellerService) *SellerHandler {
    return &SellerHandler{
        service: service,
    }
}

// ------------------------------------------------------------------------------ MAIN HANDLER
func (h *SellerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/sellers/options" {
			h.AllSellers(w, r)
			return
		}

		sellerID := r.PathValue("id")

		if sellerID != "" {
				h.getSeller(w, r, sellerID)
				return
		}

		h.listSellers(w, r)

}


// ------------------------------------------------------------------------------------ ALL SELLERS
func (h *SellerHandler) AllSellers (w http.ResponseWriter, r *http.Request) {
    sellers, err := h.service.AllSellers(r.Context())
    if err != nil {
				log.Printf("SellerList error: %v", err)

    		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(sellers)
}

// ------------------------------------------------------------------------------------------ SINGLE SELLER BY ID
func (h *SellerHandler) getSeller(w http.ResponseWriter, r *http.Request, sellerID string) {
    seller, err := h.service.SellerById(r.Context(), sellerID)
    if err != nil {
				log.Printf("SellerList error: %v", err)

    		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(seller)
}

// ------------------------------------------------------------------------------------ LIST SELLERS
func (h *SellerHandler) listSellers (w http.ResponseWriter, r *http.Request) {
    filter, err := parseSellerFilter(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    sellers, err := h.service.SellerList(r.Context(), filter)
    if err != nil {
				log.Printf("SellerList error: %v", err)

    		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(sellers)
}


func parseSellerId(r *http.Request) (*string, error) {
	var sellerId string

  q := r.URL.Query()
  sellerId = q.Get("id")
	
	if sellerId == "" {
		return nil, errors.New("id required")
	}
	return &sellerId, nil
}


func parseSellerFilter(r *http.Request) (model.SellerFilter, error) {
    var filter model.SellerFilter

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
        sort, err := parseSellerSort(v)
        if err != nil {
            return filter, err
        }

        filter.Sort = sort
    }

    return filter, nil
}


// ---------------------------------------------- parseSellerSort
func parseSellerSort(value string) ([]model.SellerSort, error) {
    parts := strings.Split(value, ",")

    result := make([]model.SellerSort, 0, len(parts))

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

        result = append(result, model.SellerSort{
            Field: field,
            Order: order,
        })
    }

    return result, nil
}
