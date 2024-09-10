package gram

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegram01/wingram-one/route"
	"github.com/codegram01/wingram-one/template"
	"github.com/go-chi/chi/v5"
)

type GramPage struct {
	template.BasePage
	Gram *Gram
}

type GramsPage struct {
	template.BasePage
	Grams []*Gram
}

func (rs *Resource) RoutesTemplate() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.ListTemplate)

	r.Route("/create", func(r chi.Router) {
		r.Get("/", rs.Template.StaticPageHandler("gram/create", "Create Gram"))
		r.Post("/", rs.CreateTemplate)
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.DetailTemplate)

		r.Route("/delete", func(r chi.Router) {
			r.Get("/", rs.GetDeleteTemplate)
			r.Post("/", rs.DeleteTemplate)
		})

		r.Route("/update", func(r chi.Router) {
			r.Get("/", rs.GetUpdateTemplate)
			r.Post("/", rs.UpdateTemplate)
		})
	})

	return r
}

func (rs *Resource) ListTemplate(w http.ResponseWriter, r *http.Request) {
	grams, err := rs.DbList()
	if err != nil {
		return
	}

	rs.Template.ServePage(w, "gram", GramsPage{
		BasePage: rs.Template.NewBasePage(r, "Grams Page"),
		Grams:    grams,
	})
}

func (rs *Resource) DetailTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := route.ReadInt(r, "id")
	if err != nil {
		return
	}

	gram, err := rs.DbDetail(id)
	if err != nil {
		return
	}

	rs.Template.ServePage(w, "gram/detail", GramPage{
		BasePage: rs.Template.NewBasePage(r, "Detail Gram Page"),
		Gram:     gram,
	})
}

func (rs *Resource) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	content := r.FormValue("content")

	gram, err := rs.DbCreate(&Gram{
		Title: title,
		Description: description,
		Content: content,
	})
	if err != nil {
		log.Println(err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/grams/%d", gram.Id), http.StatusMovedPermanently)
}

func (rs *Resource) GetDeleteTemplate(w http.ResponseWriter, r *http.Request) {
	log.Println("on get delete")
	id, err := route.ReadInt(r, "id")
	if err != nil {
		return
	}

	gram, err := rs.DbDetail(id)
	if err != nil {
		return
	}

	rs.Template.ServePage(w, "gram/delete", GramPage{
		BasePage: rs.Template.NewBasePage(r, "Delete Gram"),
		Gram:     gram,
	})
}

func (rs *Resource) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := route.ReadInt(r, "id")
	if err != nil {
		return
	}

	err = rs.DbDelete(id)
	if err != nil {
		return
	}

	http.Redirect(w, r, "/grams", http.StatusMovedPermanently)
}

func (rs *Resource) GetUpdateTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := route.ReadInt(r, "id")
	if err != nil {
		return
	}

	gram, err := rs.DbDetail(id)
	if err != nil {
		return
	}

	rs.Template.ServePage(w, "gram/update", GramPage{
		BasePage: rs.Template.NewBasePage(r, "Update Gram"),
		Gram:     gram,
	})
}

func (rs *Resource) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := route.ReadInt(r, "id")
	if err != nil {
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	content := r.FormValue("content")

	gram, err := rs.DbUpdate(id, &Gram{
		Title: title,
		Description: description,
		Content: content,
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("success update")
	http.Redirect(w, r, fmt.Sprintf("/grams/%d", gram.Id), http.StatusMovedPermanently)
}