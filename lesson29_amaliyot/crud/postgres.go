package crud

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "go"
	password = "1918"
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, dbname, password)
	Db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = Db.Ping()
	if err != nil {
		return nil, err
	}
	return Db, nil

}
