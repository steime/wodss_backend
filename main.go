package main

import (
	"fmt"
	"net/http"
	"wodss_backend/handler"
	"wodss_backend/mysql"
)

func main() {
	repository := mysql.NewMySqlRepository()
	fmt.Printf("Server started on port 8080...\n")
	http.ListenAndServe(":8080", handler.NewRouter(repository))
}

/*
func userHandler(w http.ResponseWriter, r *http.Request) {
	u := User{Id: 1, Name: "Dominik"}
	users := []User{u}
	json, _ := json.Marshal(users)
	w.Write(json)
}

func userHandler(repository data.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := repository.AllUsers()
		json, _ := json.Marshal(users)
		w.Write(json)
	}
}*/
