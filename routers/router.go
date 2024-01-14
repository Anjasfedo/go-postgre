package routers

import (
	"github.com/Anjasfedo/go-postgres/middleware"
	"github.com/gorilla/mux"
)

// Router function returns a new instance of Gorilla Mux router configured with endpoints and corresponding middleware
func Router() *mux.Router {
	// Create a new instance of Gorilla Mux router
	r := mux.NewRouter()

	// Define endpoints and their corresponding middleware handlers

	// GET request to "/api/stock", handled by middleware.GetStocks function
	r.HandleFunc("/api/stock", middleware.GetStocks).Methods("GET", "OPTIONS")

	// GET request to "/api/stock/{id}", handled by middleware.GetStockById function
	r.HandleFunc("/api/stock/{id}", middleware.GetStockById).Methods("GET", "OPTIONS")

	// POST request to "/api/stock", handled by middleware.CreateStock function
	r.HandleFunc("/api/stock", middleware.CreateStock).Methods("POST", "OPTIONS")

	// PUT request to "/api/stock/{id}", handled by middleware.UpdateStockById function
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStockById).Methods("PUT", "OPTIONS")

	// DELETE request to "/api/stock/{id}", handled by middleware.DeleteStockById function
	r.HandleFunc("/api/stock/{id}", middleware.DeleteStockById).Methods("DELETE", "OPTIONS")

	// Return the configured router
	return r
}
