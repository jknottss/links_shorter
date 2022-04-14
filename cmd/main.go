package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"new_ozon_test/internal/connection"
	"new_ozon_test/internal/handler"
	"new_ozon_test/internal/storage"
	"os"
)

func main() {
	fmt.Println("Server starting")
	fmt.Println("Storage Type -", storage.TypeStorage)
	if os.Getenv("STORAGE") == "PSQL" {
		connection.StartServer()
	}
	handler.RequestHandler()
}
