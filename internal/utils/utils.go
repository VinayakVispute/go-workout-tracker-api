package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]interface{} // interface{} is a generic type that can hold any value
// Envelope is a map that can hold any type of data. It is used to create a JSON response with a single key-value pair.

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {

	js, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ReadIDParam(r *http.Request) (int64, error) {

	paramsId := chi.URLParam(r, "id")

	if paramsId == "" {
		return 0, errors.New("invalid id parameter")
	}

	id, err := strconv.ParseInt(paramsId, 10, 64)

	if err != nil {
		return 0, errors.New("invalid id parameter type")
	}

	return id, nil
}
