package connection

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

var schema = `CREATE TABLE IF NOT EXISTS links (
		full_link	text,
		short_link	text
	);`

type Connection struct {
	Conn *sqlx.DB
}

var Con = Connection{}

func StartServer() {
	Con = OpenConnection()
}

func OpenConnection() (conf Connection) {
	config, err := os.LookupEnv("POSTGRES_URL")
	if !err {
		log.Fatal(err)
	}
	conn, OpenErr := sqlx.Connect("pgx", config)
	if OpenErr != nil {
		log.Println("database not ready yet, please wait 15 seconds")
		t := time.NewTimer(15 * time.Second)
		<-t.C
		conn, OpenErr = sqlx.Connect("pgx", config)
		if OpenErr != nil {
			log.Fatal(err)
		}
	}
	log.Println("connection successful")
	_, _ = conn.Exec(schema)
	conf.Conn = conn
	return
}
