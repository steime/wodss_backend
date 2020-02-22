package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wodss_backend/data"
)

func MakeProductsHandler(repository data.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products := repository.AllUsers()

		if json, err := json.Marshal(products); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
