package account

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/codegram01/wingram-one/crypto"
	"github.com/codegram01/wingram-one/key"
	"github.com/codegram01/wingram-one/profile"
	"github.com/codegram01/wingram-one/route"
	"github.com/codegram01/wingram-one/token"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.ListHandler)
	r.Post("/register", rs.RegisterHandler)
	r.Post("/login", rs.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(rs.AuthMiddleware)

		// get my info by token
		r.Get("/info", rs.InfoHandler)
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.DetailHandler)
	})

	return r
}

func (rs *Resource) ListHandler(w http.ResponseWriter, r *http.Request) {
	accs, err := rs.DbInfoList()

	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteJson(w, accs)
}

type RegisterReq struct {
	AccountReq
	profile.ProfileReq
}

func (rs *Resource) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var registerReq RegisterReq
	err := route.ReadJsonBody(w, r.Body, &registerReq)

	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	passHash, err := crypto.Hash(registerReq.Password)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	acc, err := rs.DbCreate(&AccountReq{
		Email:    registerReq.Email,
		Password: passHash,
	})
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	profileResource := &profile.Resource{
		Db: rs.Db,
	}
	profile, err := profileResource.DbCreate(&profile.ProfileReq{
		Name:      registerReq.Name,
		AccountId: acc.Id,
	})
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteJson(w, &AccountInfo{
		Id:        acc.Id,
		Email:     acc.Email,
		Name:      profile.Name,
		ProfileId: profile.Id,
	})
}

func (rs *Resource) DetailHandler(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	acc, err := rs.DbInfoDetail(id)
	if err != nil {
		if err == sql.ErrNoRows {
			route.WriteError(w, http.StatusNotFound, errors.New(key.ErrNotFound))
			return
		}
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	route.WriteJson(w, acc)
}

func (rs *Resource) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq AccountReq
	err := route.ReadJsonBody(w, r.Body, &loginReq)

	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	accAuth, err := rs.DbDetailAuth(loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			route.WriteError(w, http.StatusNotFound, errors.New(key.ErrNotFound))
			return
		}
		route.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if !crypto.CheckHash(loginReq.Password, accAuth.Password) {
		route.WriteError(w, http.StatusBadRequest, errors.New(key.ErrPassword))
		return
	}

	// user right
	// return token here
	tokenResource := &token.Resource{
		Db: rs.Db,
	}

	idTkn, err := tokenResource.CreateToken(int64(accAuth.Id))
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
	}

	tkn, err := crypto.GenerateToken(jwt.StandardClaims{
		Id:        strconv.FormatInt(idTkn, 10),
		ExpiresAt: time.Now().Add(time.Minute * 6000).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   "TOKEN",
	})
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
	}

	idRefreshTkn, err := tokenResource.CreateRefreshToken(token.Token{
		Token_id:   idTkn,
		Account_id: accAuth.Id,
	})
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
	}

	refreshTkn, err := crypto.GenerateToken(jwt.StandardClaims{
		Id:        strconv.FormatInt(idRefreshTkn, 10),
		ExpiresAt: time.Now().Add(time.Minute * 6000).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   "REFRESH_TOKEN",
	})
	if err != nil {
		route.WriteError(w, http.StatusInternalServerError, err)
	}

	route.WriteJson(w, &Auth{
		Access_token:  tkn,
		Refresh_token: refreshTkn,
	})
}

func (rs *Resource) InfoHandler(w http.ResponseWriter, r *http.Request) {
	identity := r.Context().Value(key.CtxIdentity).(*Identity)

	route.WriteJson(w, identity.AccountInfo)
}
