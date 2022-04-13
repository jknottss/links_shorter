package connection

import (
	"fmt"
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
		fmt.Println("this error")
		log.Fatal(err)
	}
	conn, OpenErr := sqlx.Connect("pgx", config)
	if OpenErr != nil {
		log.Println("wait 30 second for reconnect")
		time.NewTimer(30 * time.Second)
		log.Println("waiting stopped")
		conn, OpenErr = sqlx.Connect("pgx", config)
		if OpenErr != nil {
			log.Fatal(err)
		}
	}
	_, _ = conn.Exec(schema)
	conf.Conn = conn
	return
}
