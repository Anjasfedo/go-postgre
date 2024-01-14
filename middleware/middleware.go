package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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

}

func GetStockById(w http.ResponseWriter, r *http.Request) {

}

func CreateStock(w http.ResponseWriter, r *http.Request) {

}

func UpdateStockById(w http.ResponseWriter, r *http.Request) {

}

func DeleteStockById(w http.ResponseWriter, r *http.Request) {

}
