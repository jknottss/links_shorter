package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Data struct {
	FullLink  string `json:"full_link" db:"full_link"`
	ShortLink string `json:"short_link" db:"short_link"`
	Message   string `json:"message"`
}

//todo интерфейсы реализовать
//type Worker interface {
//	PasteLink(string)(string, error)
//	GetLink(string)(string, error)
//}

var longLinks map[string]string
var shortLinks map[string]string

func (s *Data) saveLongLink() (string, error) {
	if s.FullLink == "" {
		return "", errors.New("empty URL")
	}
	if shortLink, ok := longLinks[s.FullLink]; !ok {
		link := CreateLink()
		longLinks[s.FullLink] = link
		shortLinks[link] = s.FullLink
		return link, nil
	} else {
		return shortLink, nil
	}
}

func (s *Data) getLongLink() (string, error) {
	if s.ShortLink == "" {
		return "", errors.New("Empty short link")
	}
	if fullLink, ok := shortLinks[s.ShortLink]; !ok {
		return "", errors.New("Full URL does not exist")
	} else {
		return fullLink, nil
	}
}

func RequestHandler(conf Connection) {
	r := mux.NewRouter()
	r.HandleFunc("/link", conf.pasteLink).Methods("POST")
	//r.HandleFunc("/link", conf.getLink).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

//todo поставить мьютексы

func (s *Connection) pasteLink(w http.ResponseWriter, r *http.Request) {
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	if s.memoryFlag == 1 {
		link, err := data.saveLongLink()
		if err != nil {
			data.Message = "Сюда вписать ошибку"
		}
		data.ShortLink = link
	} else {
		err := s.conn.Get(&data, `SELECT * FROM db WHERE full_link=$1;`, data.FullLink)
		if err != nil {
			link := CreateLink()
			if err != nil {
				log.Fatal(err)
			}
			data.ShortLink = link
			_, err := s.conn.NamedQuery(`INSERT INTO db VALUES (:full_link, :short_link)`, data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	json.NewEncoder(w).Encode(data)
}

// func (s *Connection) getLink(w http.ResponseWriter, r *http.Request) {

// }

const allowChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func CreateLink() string {
	rand.Seed(time.Now().UnixNano())
	shortLink := make([]byte, 10)
	for i := range shortLink {
		shortLink[i] = allowChars[rand.Intn(len(allowChars))]
	}
	return string(shortLink)
}

type Connection struct {
	memoryFlag int
	conn       *sqlx.DB
}

func StartServer() {
	storage := os.Getenv("STORAGE")
	conf := Connection{}
	defer conf.conn.Close()
	if storage == "PSQL" {
		conf = OpenConnection()
	} else if storage == "INMEMORY" {
		conf.memoryFlag = 1
	}
	RequestHandler(conf)
}

func OpenConnection() (conf Connection) {
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		log.Fatal("Enter user for psql in .env")
	}
	psw, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		log.Fatal("Enter password for psql in .env")
	}
	psqlConf := fmt.Sprintf("postgres://%s:%s@postgres:5432", user, psw)
	conn, OpenErr := sqlx.Connect("pgx", psqlConf)
	if OpenErr != nil {
		log.Fatal(OpenErr)
	}
	conf.conn = conn
	return
}

func main() {
	StartServer()
}
