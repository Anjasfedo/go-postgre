# Go RESTful API with PostgreSQL 🚀

This project is a simple RESTful API built using GoLang with PostgreSQL.

## Getting Started🛠️

Follow these steps to run the project locally:

1. **Set Up Database:**

   Make sure you have PostgreSQL installed and create a database named go_postgres. Update the .env file with your PostgreSQL connection URL.

2. **Set Up Database:**

```bash
go run main.go
```

The server will start at http://localhost:8000.

## API Endpoints🚪

### Get Stocks

- **Endpoint:** `GET /api/stock`
- **Description:** Retrieve a list of all Stocks.

### Get Stock by ID

- **Endpoint:** `GET /api/stock/:id`
- **Description:** Retrieve a specific stock by ID.

### Create Stock

- **Endpoint:** `POST /api/stock`
- **Description:** Create a new stock.

### Update Stock by ID

- **Endpoint:** `PUT /api/stock/:id`
- **Description:** Update a specific stock by ID.

### Delete Stock by ID

- **Endpoint:** `DELETE /api/stock/:id`
- **Description:** Delete a specific stock by ID.

The API Postman Collection is available in the /postman-collection/ directory.

## Dependencies 📦

- [gorilla/mux](https://github.com/gorilla/mux): HTTP request router
- [pq](https://github.com/lib/pq): PostgreSQL driver for Go

## Closing Notes📝

If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

Happy coding!🚀👨‍💻
