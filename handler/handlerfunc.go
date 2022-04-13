package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"new_ozon_test/storage"
)

func (app *App) PasteLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&app.Data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	app.Data, err = app.StorType.AddLink(app.Data.FullLink)
	if err != nil {
		log.Println(err)
		w.WriteHeader(405)
	}
	_ = json.NewEncoder(w).Encode(app.Data)
	app.Data = storage.Data{}
}

func (app *App) GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&app.Data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	app.Data, err = app.StorType.GetLink(app.Data.ShortLink)
	if err != nil {
		log.Println(err)
		w.WriteHeader(405)
	}
	_ = json.NewEncoder(w).Encode(app.Data)
	app.Data = storage.Data{}
}
