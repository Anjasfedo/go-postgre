package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Anjasfedo/go-postgres/middleware"
	"github.com/Anjasfedo/go-postgres/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Postgres")

	return db
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := middleware.GetHandler()
	if err != nil {
		log.Fatalf("Unable to Get Stocks. %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Can't Convert the String into Int. %v", err)
	}

	stock, err := middleware.GetByIdHandler(int64(ID))
	if err != nil {
		log.Fatalf("Unable to Get Stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Can't Decode the Request Body. %v", err)
	}

	createdID := middleware.CreateHandler(stock)

	res := Response{
		ID:      createdID,
		Message: "Stock Created",
	}

	json.NewEncoder(w).Encode(res)

}

func UpdateStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Can't Connvert the String into Int. %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(stock)
	if err != nil {
		log.Fatalf("Can't Decode the Request Body. %v", err)
	}

	updatedStock := middleware.UpdateByIdHandler(int64(ID), stock)

	msg := fmt.Sprintf("Stock Updated. Total row affected %v", updatedStock)

	res := Response{
		ID:      int64(ID),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Can't Convert the String into Int. %v", err)
	}

	deletedStock := middleware.DeleteByIdHandler(int64(ID))

	msg := fmt.Sprintf("Stock Deleted. Total row affected %v", deletedStock)

	res := Response{
		ID:      int64(ID),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
