package models

import (
	"database/sql"
	"fmt"
)

func createSamples(db *sql.DB) {
	fmt.Println("Creating sample data..")
	db.Exec("INSERT INTO USERS(username, email, password) VALUES('usertest', 'test@email.com', '123456');")
}
