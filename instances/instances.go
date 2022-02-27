package instances

import (
	"database/sql"
)

func GetDatabaseInstance() *sql.DB {
	connStr := "postgres://dev:dev@db:5432/dev?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	return db
}
