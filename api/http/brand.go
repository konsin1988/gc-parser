package http

import (
  "encoding/json"
  "net/http"
	"context"
	_ "errors"
	"log"

	"konsin1988/gc-api/model"
)

type BrandService interface {
		AllBrands(ctx context.Context) ([]model.AllBrandsItem, error)
		BrandList(ctx context.Context, filter model.Filter) ([]model.BrandListItem, error)
		BrandById(ctx context.Context, BrandId string) (*model.ResponseBrandById, error)
}

type BrandHandler struct {
    service BrandService
}

func NewBrandHandler(service BrandService) *BrandHandler {
    return &BrandHandler{
        service: service,
    }
}

// ------------------------------------------------------------------------------ MAIN HANDLER
func (h *BrandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/brands/options" {
			h.AllBrands(w, r)
			return
		}

		brandID := r.PathValue("id")

		if brandID != "" {
				h.getBrand(w, r, brandID)
				return
		}

		h.listBrands(w, r)

}


// ------------------------------------------------------------------------------------ ALL BRANDS 
func (h *BrandHandler) AllBrands (w http.ResponseWriter, r *http.Request) {
    brands, err := h.service.AllBrands(r.Context())
    if err != nil {
				log.Printf("BrandList error: %v", err)

    		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(brands)
}


// ------------------------------------------------------------------------------------------ SINGLE BRAND BY ID
func (h *BrandHandler) getBrand(w http.ResponseWriter, r *http.Request, brandID string) {
    brand, err := h.service.BrandById(r.Context(), brandID)
    if err != nil {
				log.Printf("BrandList error: %v", err)

    		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(brand)
}

// ------------------------------------------------------------------------------------------- LIST BRANDS
func (h *BrandHandler) listBrands(w http.ResponseWriter, r *http.Request) {
    filter, err := parseFilter(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
		brands, err := h.service.BrandList(r.Context(), filter)
    if err != nil {
				log.Printf("BrandList error: %v", err)

    		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(brands)
}
