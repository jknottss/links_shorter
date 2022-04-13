package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

//todo поставить мьютексы
//todo установить статусы ответа

func (app *App) PasteLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&app.data)
	if err != nil {
		log.Println(err)
	}
	app.data, err = app.storType.AddLink(app.data.FullLink)
	if err != nil {
		log.Println(err)
		//установить статус ошибочный
	}
	_ = json.NewEncoder(w).Encode(app.data)
}

func (app *App) GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&app.data)
	if err != nil {
		log.Println(err)
	}
	app.data, err = app.storType.GetLink(app.data.ShortLink)
	if err != nil {
		log.Println(err)
		//установить статус ошибочный
	}
	_ = json.NewEncoder(w).Encode(app.data)
}
