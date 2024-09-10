package post

import (
	"net/http"

	"github.com/codegram01/wingram-one/account"
	"github.com/codegram01/wingram-one/key"
	"github.com/codegram01/wingram-one/route"
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
		return
	}

	rs.Template.ServePage(w, "post", PostsPage{
		BasePage: rs.Template.NewBasePage(r, "Posts Page"),
		Posts: posts,
	})
}

func (rs *Resource) DetailTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := route.ReadInt(r, "id")
	if err != nil {
		return
	}

	post, err := rs.DbDetail(id)
	if err != nil {
		return
	}

	rs.Template.ServePage(w, "post/detail", PostPage{
		BasePage: rs.Template.NewBasePage(r, "Detail Post Page"),
		Post: post,
	})
}

func (rs *Resource) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	identity := r.Context().Value(key.CtxIdentity).(*account.Identity)

	var postReq Post
	err := route.ReadJsonBody(w, r.Body, &postReq)

	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	postReq.ProfileId = identity.ProfileId

	post, err := rs.DbCreate(&postReq)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteJson(w, post)
}
