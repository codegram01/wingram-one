package gram

import (
	"testing"

	"github.com/codegram01/wingram-one/config"
	"github.com/codegram01/wingram-one/database"
)

func TestDbCreate(t *testing.T) {
	cfg := config.Init()

	db, err := database.Connect(cfg)
	if err != nil {
		t.Fatal(err)
	}

	rs := &Resource{
		Db: db,
	}

	g, err := rs.DbCreate(&Gram{
		Title: "hello 2",
		Description: "Hello World 2",
		Content: "Welcome to my home page 2",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(g)
}