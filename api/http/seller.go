package http

import (
  "encoding/json"
  "net/http"
	"context"
	"strconv"
	"fmt"
	"log"

	"konsin1988/gc-api/model"
)

type SellerService interface {
    SellerList(ctx context.Context, filter model.SellerFilter) ([]model.ResponseSellerListItem, error)
}

type SellerHandler struct {
    service SellerService
}

func NewSellerHandler(service SellerService) *SellerHandler {
    return &SellerHandler{
        service: service,
    }
}


func (h *SellerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

    return filter, nil
}
