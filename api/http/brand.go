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
    //SellerList(ctx context.Context, filter model.Filter) ([]model.SellerListItem, error)
		//SellerById(ctx context.Context, SellerId string) (*model.ResponseSellerById, error)
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

		//brandID := r.PathValue("id")

		//if brandID != "" {
		//		h.getSeller(w, r, brandID)
		//		return
		//}

		//h.listBrands(w, r)

}


// ------------------------------------------------------------------------------------ ALL BRANDS 
func (h *BrandHandler) AllBrands (w http.ResponseWriter, r *http.Request) {
    brands, err := h.service.AllBrands(r.Context())
    if err != nil {
				log.Printf("SellerList error: %v", err)

    		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(brands)
}
