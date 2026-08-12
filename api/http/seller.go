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
		sellerID := r.PathValue("id")

		if sellerID != "" {
				h.getSeller(w, r, sellerID)
				return
		}

		h.listSellers(w, r)

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

    if v := q.Get("brand"); v != "" {
        id, err := strconv.Atoi(v)
        if err != nil {
            return filter, fmt.Errorf("invalid brand")
        }
        filter.BrandID = &id
    }

    if v := q.Get("category"); v != "" {
        id, err := strconv.Atoi(v)
        if err != nil {
            return filter, fmt.Errorf("invalid category")
        }
        filter.CategoryID = &id
    }

    if v := q.Get("min_goods"); v != "" {
        n, err := strconv.Atoi(v)
        if err != nil {
            return filter, fmt.Errorf("invalid minGoods")
        }
        filter.MinGoods = &n
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
