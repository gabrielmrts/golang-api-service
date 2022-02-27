package models

import (
	"github.com/gabrielmrts/golang-api-service/instances"
	_ "github.com/lib/pq"
)

func Init() {
	db := instances.GetDatabaseInstance()
	createTables(db)
	createSamples(db)
}
