package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	stocks, err := getHandler()
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

	stock, err := getByIdHandler(int64(ID))
	if err != nil {
		log.Fatalf("Unable to Get Stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Can't Decode the Request Body. %v", err)
	}

	createdID := createHandler(stock)

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

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Can't Decode the Request Body. %v", err)
	}

	updatedStock := updateByIdHandler(int64(ID), stock)

	msg := fmt.Sprintf("Stock Updated. Total Row affected %v", updatedStock)

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

	deletedStock := deleteByIdHandler(int64(ID))

	msg := fmt.Sprintf("Stock Deleted. Total Row affected %v", deletedStock)

	res := Response{
		ID:      int64(ID),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

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

func getByIdHandler(ID int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`

	row := db.QueryRow(sqlStatement, ID)

	err := row.Scan(&stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No Rows Return!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Can't Scan the Row. %v", err)
	}

	return stock, err
}

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
