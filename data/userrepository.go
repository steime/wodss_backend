package data

type Repository interface {
	//AddUser(u User)
	AllUsers() []User
}

type DefaultRepository struct {
	Users []User
}

/*
func (r *DefaultRepository) AddUser(u User) {
	r.Users = append(r.Users, u)
}*/

func (r *DefaultRepository) AllUsers() []User {
	return r.Users
}
