package server

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/codegram01/wingram-one/config"
	"github.com/codegram01/wingram-one/database"
	"github.com/codegram01/wingram-one/template"
	"github.com/codegram01/wingram-one/template/templates"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	mode string
	template *template.Template
	mux       *chi.Mux
	db *database.Db
	staticFS  fs.FS
}

type ServerCfg struct {
	Cfg *config.Config
	Db  *database.Db
}

func Init(scfg *ServerCfg) {
	ts, err := templates.ParsePageTemplates()
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}

	temp := &template.Template{
		Templates: ts,
	}

	r := chi.NewRouter()

	server := &Server{
		mode: scfg.Cfg.Mode,
		mux:       r,
		template: temp,
		db: scfg.Db,
		staticFS:  os.DirFS("static/public"),
	}

	server.MakeHandler()

	portServer := ":" + scfg.Cfg.Port
	log.Printf("Server running %s\n", portServer)
	log.Fatal(http.ListenAndServe(portServer, r))
}
