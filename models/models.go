package models

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/gabrielmrts/golang-api-service/factories"
	_ "github.com/lib/pq"
)

func Init() {
	db := factories.GetDatabaseInstance()

	fmt.Println("Creating tables..")
	createTables(db)
	createSamples(db)
}

func createTables(db *sql.DB) {
	files, err := ioutil.ReadDir("./database")

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
