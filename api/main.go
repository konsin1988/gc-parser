package main

import (
  "net/http"
  "log"

  config "konsin1988/gc-api/config"
  repo "konsin1988/gc-api/repository"
	"konsin1988/gc-api/service"
	http_t "konsin1988/gc-api/http"
)

func main() {
	config.ConnectDB()
  defer config.DB.Close()

	repo.RunMigrations()
  repo := repo.NewRepository(config.DB)

	healthService := service.NewHealthService(repo)
  healthHandler := http_t.NewHealthHandler(healthService)

	sellerService := service.NewSellerService(repo)
	sellerHandler := http_t.NewSellerHandler(sellerService)

	brandService := service.NewBrandService(repo)
	brandHandler := http_t.NewBrandHandler(brandService)

	api := http.NewServeMux()
	
	api.Handle("GET /health", healthHandler)

	api.HandleFunc("GET /sellers/", func(w http.ResponseWriter, r *http.Request) {
	    http.Redirect(w, r, "/sellers", http.StatusMovedPermanently)
	})

	// sellers
	api.Handle("GET /sellers", sellerHandler)
	api.Handle("GET /seller/{id}", sellerHandler)
	api.Handle("GET /sellers/options", sellerHandler)

	// brands
	api.Handle("GET /brands", brandHandler)
	api.Handle("GET /brand/{id}", brandHandler)
	api.Handle("GET /brands/options", brandHandler)

	//root.Handle("/api/", auth.Middleware(api))

  log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", api))
}

