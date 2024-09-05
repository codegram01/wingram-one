package database

import (
	"database/sql"
	"log"

	"github.com/codegram01/wingram-one/config"
	_ "github.com/lib/pq"
)

type Db struct {
	Con *sql.DB // Database connection
}

func Connect(cfg *config.Config) (*Db, error) {
	con, err := sql.Open("postgres", cfg.GetDbConStr())
	if err != nil {
		return nil, err
	}

	pingErr := con.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	log.Println("DB Connected!")

	db := &Db{
		Con: con,
	}

	return db, nil
}
