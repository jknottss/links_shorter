package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"new_ozon_test/connection"
	"new_ozon_test/handler"
	"os"
)

//todo интерфейсы реализовать
//
//type Storage interface {
//	PasteLink(string) (string, error)
//	GetLink(string) (string, error)
//}

//todo настроить передачу параметров для запуска
//todo юнит-тесты
//todo не запускать постгрес если нет необходимости

func main() {
	_ = os.Setenv("STORAGE", "1")
	fmt.Println("Server start")
	flag := connection.StartServer()
	if flag == 1 {
		handler.RequestHandler()
	}
}
