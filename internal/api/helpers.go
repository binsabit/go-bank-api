package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func ReadJSON(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func GetFromParams(target string, r *http.Request) string {
	params := mux.Vars(r)
	res := params[target]
	return res
}
