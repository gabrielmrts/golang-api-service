package models

import (
	"github.com/gabrielmrts/golang-api-service/instances"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (s User) Create(user *User) int64 {
	db := instances.GetDatabaseInstance()

	stmt, err := db.Prepare("INSERT INTO users(username, email, password) VALUES($1, $2, $3)")

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(user.Username, user.Email, user.Password)

	if err != nil {
		//Need replace it to return http response
		panic(err)
	}

	affected, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	return affected
}
