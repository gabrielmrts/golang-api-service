package models

import (
	"github.com/gabrielmrts/golang-api-service/instances"
)

type User struct {
	Username string
	Email    string
	Password string
}

func (s User) Create(user *User) int64 {
	db := instances.GetDatabaseInstance()

	stmt, err := db.Prepare("INSERT INTO users(username, email, password) VALUES($1, $2, $3)")

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(user.Username, user.Email, user.Password)

	if err != nil {
		panic(err)
	}

	affected, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	return affected
}
