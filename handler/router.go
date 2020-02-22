package handler

import (
	"wodss_backend/data"

	"github.com/gorilla/mux"
)

func NewRouter(repository data.Repository) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", MakeProductsHandler(repository)).Methods("GET")
	/*r.HandleFunc("/catalog/products", MakeAddProductHandler(repository)).Methods("POST")
	r.HandleFunc("/catalog/products/{id}", MakeGetProductHandler(repository)).Methods("GET")
	r.HandleFunc("/catalog/products/{id}", MakeDeleteProductHandler(repository)).Methods("DELETE")
	r.HandleFunc("/catalog/products/{id}", MakeUpdateProductHandler(repository)).Methods("PUT")
	*/
	return r
}
