package main

import (
	"log"

	"github.com/codegram01/wingram-one/config"
	"github.com/codegram01/wingram-one/database"
	"github.com/codegram01/wingram-one/server"
)

func main() {
	cfg := config.Init()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server.Init(&server.ServerCfg{
		Cfg: cfg,
		Db:  db,
	})
}
