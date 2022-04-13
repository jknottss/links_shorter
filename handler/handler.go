package handler

import (
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"new_ozon_test/connection"
	"new_ozon_test/storage"
	"os"
	"sync"
)

type App struct {
	Data     storage.Data
	StorType storage.Storage
}

func initApp() (*App, error) {
	app := App{}
	if storage.TypeStorage == "PSQL" {
		app.StorType = &storage.Psql{
			Conn: connection.Con.Conn,
			Mu:   new(sync.Mutex),
		}
		return &app, nil
	} else if storage.TypeStorage == "INMEMORY" {
		app.StorType = &storage.Memory{
			LongLinks:  make(map[string]string),
			ShortLinks: make(map[string]string),
			Mu:         new(sync.Mutex),
		}
		return &app, nil
	}
	return &app, errors.New("wrong storage")
}
func RequestHandler() {
	app, err := initApp()
	if err != nil {
		log.Println(err)
	}
	if os.Getenv("STORAGE") == "PSQL" {
		defer connection.Con.Conn.Close()
	}
	r := mux.NewRouter()
	r.HandleFunc("/link", app.PasteLink).Methods("POST")
	r.HandleFunc("/link", app.GetLink).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
