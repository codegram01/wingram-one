package post

import (
	"net/http"
	"strconv"

	"github.com/codegram01/wingram-one/template"
	"github.com/go-chi/chi/v5"
)

type PostPage struct {
	template.BasePage
	Post *Post
}

type PostsPage struct {
	template.BasePage
	Posts []*Post
}

func (rs *Resource) RoutesTemplate() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.ListTemplate)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.DetailTemplate)
	})

	return r
}

func (rs *Resource) ListTemplate(w http.ResponseWriter, r *http.Request) {
	posts, err := rs.DbList()
	if err != nil {
		// route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	rs.Template.ServePage(w, "post", PostsPage{
		BasePage: rs.Template.NewBasePage(r, "Posts Page"),
		Posts: posts,
	})
}

func (rs *Resource) DetailTemplate(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		// route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	post, err := rs.DbDetail(id)
	if err != nil {
		// route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	rs.Template.ServePage(w, "post/detail", PostPage{
		BasePage: rs.Template.NewBasePage(r, "Detail Post Page"),
		Post: post,
	})
}
