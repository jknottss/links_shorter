package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"new_ozon_test/connection"
	"new_ozon_test/handler"
	"new_ozon_test/storage"
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
