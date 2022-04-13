package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"new_ozon_test/storage"
)

func (app *App) PasteLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&app.data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	app.data, err = app.storType.AddLink(app.data.FullLink)
	if err != nil {
		log.Println(err)
		w.WriteHeader(405)
	}
	_ = json.NewEncoder(w).Encode(app.data)
	app.data = storage.Data{}
}

func (app *App) GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&app.data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	app.data, err = app.storType.GetLink(app.data.ShortLink)
	if err != nil {
		log.Println(err)
		w.WriteHeader(405)
	}
	_ = json.NewEncoder(w).Encode(app.data)
	app.data = storage.Data{}
}
