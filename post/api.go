package post

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/codegram01/wingram-one/account"
	"github.com/codegram01/wingram-one/key"
	"github.com/codegram01/wingram-one/route"
	"github.com/go-chi/chi/v5"
)

func (rs *Resource) RoutesApi() chi.Router {
	accResource := &account.Resource{
		Db: rs.Db,
	}

	r := chi.NewRouter()

	r.Get("/", rs.ListApi)

	r.Group(func(r chi.Router) {
		r.Use(accResource.AuthMiddleware)
		r.Post("/", rs.CreateApi)
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.DetailApi)

		r.Group(func(r chi.Router) {
			r.Use(accResource.AuthMiddleware)

			r.Patch("/", rs.UpdateApi)
			r.Delete("/", rs.DeleteApi)
		})
	})

	return r
}

func (rs *Resource) ListApi(w http.ResponseWriter, r *http.Request) {
	posts, err := rs.DbList()
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteJson(w, posts)
}

func (rs *Resource) DetailApi(w http.ResponseWriter, r *http.Request) {
	id, err := route.ReadInt(r, "id")
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	post, err := rs.DbDetail(id)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteJson(w, post)
}

func (rs *Resource) CreateApi(w http.ResponseWriter, r *http.Request) {
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

func (rs *Resource) UpdateApi(w http.ResponseWriter, r *http.Request) {
	identity := r.Context().Value(key.CtxIdentity).(*account.Identity)

	idS := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var postReq Post
	err = route.ReadJsonBody(w, r.Body, &postReq)

	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	post, err := rs.DbDetail(id)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if post.ProfileId != identity.ProfileId {
		route.WriteError(w, http.StatusForbidden, errors.New(key.ErrNotHavePermission))
		return
	}

	postUpdate, err := rs.DbUpdate(id, &postReq)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteJson(w, postUpdate)
}

func (rs *Resource) DeleteApi(w http.ResponseWriter, r *http.Request) {
	identity := r.Context().Value(key.CtxIdentity).(*account.Identity)

	idS := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	post, err := rs.DbDetail(id)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if post.ProfileId != identity.ProfileId {
		route.WriteError(w, http.StatusForbidden, errors.New(key.ErrNotHavePermission))
		return
	}

	err = rs.DbDelete(id)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteSuccess(w)
}
