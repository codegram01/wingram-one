package server

import (
	"log"
	"net/http"

	"github.com/codegram01/wingram-one/account"
	"github.com/codegram01/wingram-one/config"
	"github.com/codegram01/wingram-one/database"
	"github.com/codegram01/wingram-one/middleware"
	"github.com/codegram01/wingram-one/post"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type ServerCfg struct {
	Cfg *config.Config
	Db  *database.Db
}

func Init(scfg *ServerCfg) {
	accountResource := &account.Resource{
		Db: *scfg.Db,
	}
	postResource := &post.Resource{
		Db: *scfg.Db,
	}

	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(middleware.CORS)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server Work"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.JsonApi)

		r.Mount("/accounts", accountResource.Routes())
		r.Mount("/posts", postResource.Routes())
	})

	portServer := ":" + scfg.Cfg.Port
	log.Printf("Server running %s\n", portServer)
	log.Fatal(http.ListenAndServe(portServer, r))
}
