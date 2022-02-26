package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Init() {
	//JUST TESTING THE DATABASE, IGNORE
	connStr := "postgres://dev:dev@db:5432/dev?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	result, err := db.Exec("CREATE TABLE TESTE (id int)")

	if err != nil {
		panic(err)
	}

	fmt.Print(result)
}
