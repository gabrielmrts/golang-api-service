package models

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

func createTables(db *sql.DB) {
	fmt.Println("Reading database files..")

	files, err := ioutil.ReadDir("./database")

	fmt.Println("Creating tables..")

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := fmt.Sprintf("./database/%s", file.Name())

		sql, err := ioutil.ReadFile(filePath)

		if err != nil {
			panic(err)
		}

		db.Exec(string(sql))
	}
}
