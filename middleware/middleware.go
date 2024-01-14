package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Anjasfedo/go-postgres/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Response struct represents the structure for API responses
type Response struct {
	ID      int64  `json:"stockid,omitempty"`
	Message string `json:"message,omitempty"`
}

// CreateConnection function establishes a connection to the PostgreSQL database
func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to Postgres")

	return db
}

// GetStocks handles the HTTP GET request to retrieve all stocks
func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getHandler()
	if err != nil {
		http.Error(w, "Unable to get stocks", http.StatusInternalServerError)
		log.Printf("Unable to Get Stocks. %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stocks)
}

// GetStockById handles the HTTP GET request to retrieve a stock by ID
func GetStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Can't Convert the String into Int. %v", err)
		return
	}

	stock, err := getByIdHandler(int64(ID))
	if err != nil {
		http.Error(w, "Unable to get stock", http.StatusInternalServerError)
		log.Printf("Unable to Get Stock. %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stock)
}

// CreateStock handles the HTTP POST request to create a new stock
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Can't decode the request body", http.StatusBadRequest)
		log.Printf("Can't Decode the Request Body. %v", err)
		return
	}

	createdID := createHandler(stock)

	res := Response{
		ID:      createdID,
		Message: "Stock Created",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// UpdateStockById handles the HTTP PUT request to update a stock by ID
func UpdateStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Can't Convert the String into Int. %v", err)
		return
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Can't decode the request body", http.StatusBadRequest)
		log.Printf("Can't Decode the Request Body. %v", err)
		return
	}

	updatedStock := updateByIdHandler(int64(ID), stock)

	msg := fmt.Sprintf("Stock Updated. Total Row affected %v", updatedStock)

	res := Response{
		ID:      int64(ID),
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// DeleteStockById handles the HTTP DELETE request to delete a stock by ID
func DeleteStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Can't Convert the String into Int. %v", err)
		return
	}

	deletedStock := deleteByIdHandler(int64(ID))

	msg := fmt.Sprintf("Stock Deleted. Total Row affected %v", deletedStock)

	res := Response{
		ID:      int64(ID),
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// getHandler retrieves all stocks from the database
func getHandler() ([]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stocks []models.Stock

	sqlStatement := `SELECT * FROM stocks`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Can't Execute the Query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Can't Scan the Row %v", err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, err
}

// getByIdHandler retrieves a stock by ID from the database
func getByIdHandler(ID int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`

	err := db.QueryRow(sqlStatement, ID).Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No Rows Return!")
	case err != nil:
		log.Fatalf("Can't Scan the Row. %v", err)
	}

	return stock, err
}

// createHandler creates a new stock in the database
func createHandler(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var ID int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&ID)
	if err != nil {
		log.Fatalf("Can't Execute the Query. %v", err)
	}

	fmt.Printf("Created a Single Record %v", ID)
	return ID
}

// updateByIdHandler updates a stock by ID in the database
func updateByIdHandler(ID int64, stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, ID, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Can't Execute the Query %v", err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error While Check the Affected Rows. %v", err)
	}

	fmt.Printf("Total Rows Affected %v", rowAffected)

	return rowAffected
}

// deleteByIdHandler deletes a stock by ID from the database
func deleteByIdHandler(ID int64) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, ID)
	if err != nil {
		log.Fatalf("Can't Execute the Query %v", err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error While Check the Affected Rows. %v", err)
	}

	fmt.Printf("Total Rows Affected %v", rowAffected)
	return rowAffected
}
