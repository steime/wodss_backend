package main

import (
	"encoding/json"
	"net/http"
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
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func (r *DefaultRepository) AddUser(u User) {
	r.Users = append(r.Users, u)
}

func (r *DefaultRepository) AllUsers() []User {
	return r.Users
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	u := User{Id: 1, Name: "Dominik"}
	users := []User{u}
	json, _ := json.Marshal(users)
	w.Write(json)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World/Index"))
}
