package mysql

import (
	"database/sql"
	"fmt"
	"wodss_backend/data"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
	db        *sql.DB
	Connected bool
}

func NewMySqlRepository() *MySqlRepository {
	r := MySqlRepository{}
	r.connect("steime", "steime", "user")
	return &r
}

func (r *MySqlRepository) connect(user string, password string, databaseName string) {
	dataSourceName := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, databaseName)
	var err error
	if r.db, err = sql.Open("mysql", dataSourceName); err != nil {
		panic(err)
	}

	if err = r.db.Ping(); err != nil {
		panic(err)
	}

	r.Connected = true
}

/*
func (r *MySqlRepository) AddUser(u data.User) {

}
*/
func (r *MySqlRepository) AllUsers() []data.User {
	rows, err := r.db.Query("select ID, name")
	if err != nil {
		panic(err)
	}

	var users []data.User
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		} else {
			users = append(users, data.User{ID: id, Name: name})
		}
	}

	return users
}

/*
func (r *MySqlRepository) FindProducts(q string) []catalog.Product {
	return []catalog.Product{}
}

func (r *MySqlRepository) ProductById(id int) (catalog.Product, bool) {
	return catalog.Product{}, false
}

func (r *MySqlRepository) Delete(id int) bool {
	return false
}

func (r *MySqlRepository) Contains(name string) bool {
	return false
}

func (r *MySqlRepository) Update(id int, p catalog.Product) bool {
	return false
}
*/
