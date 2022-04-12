package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"new_ozon_test/connection"
	"new_ozon_test/createlink"
	"new_ozon_test/storage"
)

//todo поставить мьютексы
//todo установить статусы ответа

func PasteLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data storage.Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	if data.FullLink == "" {
		log.Print("empty URL")
		return
	}
	if connection.Conf.MemoryFlag == 1 {
		link, err := data.SaveLongLink()
		if err != nil {
			log.Println(err)
		}
		data.ShortLink = link
	} else {
		err := connection.Conf.Conn.Get(&data, "SELECT * FROM links WHERE full_link=$1;", data.FullLink)
		if err != nil {
			link := createlink.CreateLink()
			if err != nil {
				log.Println(err)
			}
			data.ShortLink = link
			_, err := connection.Conf.Conn.NamedQuery("INSERT INTO links VALUES (:full_link, :short_link)", data)
			if err != nil {
				log.Println(err)
			}
		}
	}
	_ = json.NewEncoder(w).Encode(data)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data storage.Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	if data.ShortLink == "" {
		log.Println("empty URL")
		return
	}
	if connection.Conf.MemoryFlag == 1 {
		link, err := data.GetLongLink()
		if err != nil {
			fmt.Println(err)
		}
		data.FullLink = link
	} else {
		err := connection.Conf.Conn.Get(&data, "SELECT * FROM links WHERE short_link=$1;", data.ShortLink)
		if err != nil {
			log.Println("url does not Exist")
			return
		}
	}
	_ = json.NewEncoder(w).Encode(data)
}
