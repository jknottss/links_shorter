package connection

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var schema = `CREATE TABLE IF NOT EXISTS links (
		full_link	varchar(80),
		short_link	varchar(80)
	);`

type Connection struct {
	MemoryFlag int
	Conn       *sqlx.DB
}

var Conf = Connection{}

func StartServer() int {
	storage := os.Getenv("STORAGE")
	if storage == "2" {
		Conf = OpenConnection()
	} else if storage == "1" {
		Conf.MemoryFlag = 1
	}
	return 1
}

func OpenConnection() (conf Connection) {
	config, err := os.LookupEnv("POSTGRES_URL")
	if !err {
		log.Fatal(err)
	}
	conn, OpenErr := sqlx.Connect("pgx", config)
	if OpenErr != nil {
		log.Fatal(OpenErr)
	}
	_, _ = conn.Exec(schema)
	conf.Conn = conn
	return
}
