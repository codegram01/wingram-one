package route

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ReadJsonBody(w http.ResponseWriter, r io.Reader, data interface{}) error {
	err := json.NewDecoder(r).Decode(&data)

	// we will add more logic validation body data here

	if err != nil {
		return err
	}

	return nil
}

func ReadQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func ReadInt(r *http.Request, key string) (int64, error) {
	idS := chi.URLParam(r, key)
	return strconv.ParseInt(idS, 10, 64)
}