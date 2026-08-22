package http

import (
  "encoding/json"
  "net/http"
	"context"
	"errors"
	"log"

	"konsin1988/gc-api/model"
)

type SellerService interface {
		AllSellers(ctx context.Context) ([]model.AllSellersItem, error)
    SellerList(ctx context.Context, filter model.Filter) ([]model.SellerListItem, error)
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
    filter, err := parseFilter(r)
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
