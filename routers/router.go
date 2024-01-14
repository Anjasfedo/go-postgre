package routers

import (
	"github.com/Anjasfedo/go-postgres/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/stock", middleware.GetStocks).Methods("GET", "OPTIONS")

	r.HandleFunc("/api/stock/{id}", middleware.GetStockById).Methods("GET", "OPTIONS")

	r.HandleFunc("/api/stock", middleware.CreateStock).Methods("POST", "OPTIONS")

	r.HandleFunc("/api/stock/{id}", middleware.UpdateStockById).Methods("PUT", "OPTIONS")

	r.HandleFunc("/api/stock/{id}", middleware.DeleteStockById).Methods("DELETE", "OPTIONS")
}
