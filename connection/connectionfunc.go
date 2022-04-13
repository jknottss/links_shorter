package connection

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

var schema = `CREATE TABLE IF NOT EXISTS links (
		full_link	varchar(80),
		short_link	varchar(80)
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
		log.Println("database not ready yet, wait 10 seconds")
		t := time.NewTimer(10 * time.Second)
		<-t.C
		conn, OpenErr = sqlx.Connect("pgx", config)
		if OpenErr != nil {
			log.Fatal(err)
		}
	}
	log.Println("reconnection successful")
	_, _ = conn.Exec(schema)
	conf.Conn = conn
	return
}
