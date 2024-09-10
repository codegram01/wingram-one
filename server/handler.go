package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegram01/wingram-one/account"
	"github.com/codegram01/wingram-one/gram"
	"github.com/codegram01/wingram-one/middleware"
	"github.com/codegram01/wingram-one/post"
	"github.com/codegram01/wingram-one/template"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) MakeHandler() {
	r := s.mux

	accountResource := &account.Resource{
		Db: s.db,
		Template: s.template,
	}
	postResource := &post.Resource{
		Db: s.db,
		Template: s.template,
	}
	gramResource := &gram.Resource{
		Db: s.db,
		Template: s.template,
	}

	r.Use(chiMiddleware.Logger)
	
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	
	// ssr templates handler
	r.Route("/", func(r chi.Router) {
		r.Get("/", s.homeHandler)
		r.Get("/test", s.testHandler)
		r.Get("/about", s.template.StaticPageHandler("about", "About Page"))

		// r.Mount("/accounts", accountResource.Routes())
		r.Mount("/posts", postResource.RoutesTemplate())
		r.Mount("/grams", gramResource.RoutesTemplate())
	})

	
	// api json handler 
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.JsonApi)

		r.Mount("/accounts", accountResource.Routes())
		r.Mount("/posts", postResource.RoutesApi())
	})

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static/public"))
	FileServer(r, "/public", filesDir)
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	s.template.ServePage(w, "home", template.BasePage{
		HTMLTitle: "Home Page",
	})
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/grams", http.StatusMovedPermanently)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}