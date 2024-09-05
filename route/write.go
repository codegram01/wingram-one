package route

import (
	"encoding/json"
	"net/http"
)

type ResJson struct {
	Status  string `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func WriteJson(w http.ResponseWriter, dataBody any) error {
	resJson := &ResJson{
		Status: "ok",
		Data:   dataBody,
	}

	err := json.NewEncoder(w).Encode(resJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}

func WriteSuccess(w http.ResponseWriter) {
	resJson := &ResJson{
		Status:  "ok",
		Message: "success",
	}
	WriteJson(w, resJson)
}

func WriteError(w http.ResponseWriter, statusCode int, errMess error) error {
	w.WriteHeader(statusCode)

	resJson := &ResJson{
		Status: "error",
		Error:  errMess.Error(),
	}

	err := json.NewEncoder(w).Encode(resJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}
