package models

import (
	"github.com/gabrielmrts/golang-api-service/instances"
)

func Init() {
	db := instances.GetDatabaseInstance()
	createTables(db)
	createSamples(db)
}
