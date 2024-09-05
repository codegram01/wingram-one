package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/codegram01/wingram-one/crypto"
	"github.com/codegram01/wingram-one/key"
	"github.com/codegram01/wingram-one/route"
)

type Auth struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type Identity struct {
	AccountInfo
}

func (rs *Resource) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			route.WriteError(w, http.StatusUnauthorized, errors.New(key.ErrAuthHeaderMiss))
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			route.WriteError(w, http.StatusUnauthorized, errors.New(key.ErrAuthHeaderInvalid))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		indentity, err := rs.IdentityToken(token)

		if err != nil {
			route.WriteError(w, http.StatusUnauthorized, errors.New(key.ErrAuthTokenInValid))
			return
		}

		ctx := context.WithValue(r.Context(), key.CtxIdentity, indentity)
		ctx = context.WithValue(ctx, key.CtxAuthorization, authHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs *Resource) IdentityToken(token string) (*Identity, error) {
	var identity Identity

	claims, err := crypto.ParseToken(token)
	if err != nil {
		return &identity, err
	}

	idToken, err := strconv.ParseInt(claims.Id, 10, 64)
	if err != nil {
		return &identity, err
	}

	idAccount, err := rs.DbGetAccountByToken(idToken)
	if err != nil {
		return &identity, err
	}

	acc, err := rs.DbInfoDetail(idAccount)
	if err != nil {
		return &identity, err
	}

	identity.AccountInfo = *acc

	return &identity, nil
}
