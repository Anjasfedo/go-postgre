package models

// Stock struct represents the data model for stock information
type Stock struct {
	StockID int64  `json:"stockid"` // StockID field with json tag "stockid"
	Name    string `json:"name"`    // Name field with json tag "name"
	Price   int64  `json:"price"`   // Price field with json tag "price"
	Company string `json:"company"` // Company field with json tag "company"
}
