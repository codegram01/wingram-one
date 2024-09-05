package database

import (
	"testing"

	"github.com/codegram01/wingram-one/config"
)

func TestConnect(t *testing.T) {

	cfg := config.Init()
	db, err := Connect(cfg)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(db)
}
