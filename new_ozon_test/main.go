package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"new_ozon_test/connection"
	"new_ozon_test/handler"
	"os"
)

//todo настроить передачу параметров для запуска
//todo юнит-тесты
//todo не запускать постгрес если нет необходимости

func main() {
	_ = os.Setenv("STORAGE", "PSQL")
	fmt.Println("Server start")
	flag := connection.StartServer()
	if flag == 1 {
		handler.RequestHandler()
	}
}
