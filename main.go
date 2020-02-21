package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int
	Name string
}

type Repository interface {
	AddUser(u User)
	AllUsers() []User
}

type DefaultRepository struct {
	Users []User
}

func main() {
	r := DefaultRepository{}
	router := mux.NewRouter()
	router.HandleFunc("/user", userHandler(&r)).Methods("GET")
	router.HandleFunc("/", indexHandler).Methods("GET")
	fmt.Printf("Server listening at :8080")
	http.ListenAndServe(":8080", router)
}

func (r *DefaultRepository) AddUser(u User) {
	r.Users = append(r.Users, u)
}

func (r *DefaultRepository) AllUsers() []User {
	return r.Users
}

/*
func userHandler(w http.ResponseWriter, r *http.Request) {
	u := User{Id: 1, Name: "Dominik"}
	users := []User{u}
	json, _ := json.Marshal(users)
	w.Write(json)
}*/

func userHandler(repository Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := repository.AllUsers()
		json, _ := json.Marshal(users)
		w.Write(json)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World/Index"))
}
