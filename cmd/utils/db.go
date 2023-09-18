package utils

import (
	// "github.com/panyam/goutils/utils"
	"log"
	"strings"

	"gorm.io/driver/postgres"

	// "gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func OpenDB(db_endpoint string) (db *gorm.DB, err error) {
	log.Println("Connecting to DB: ", db_endpoint)
	if strings.HasPrefix(db_endpoint, "postgres://") {
		db, err = gorm.Open(postgres.Open(db_endpoint), &gorm.Config{})
		/*
			} else if strings.HasPrefix(db_endpoint, "sqlite://") {
				dbpath := utils.ExpandUserPath((db_endpoint)[len("sqlite://"):])
				db, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
		*/
	}
	if err != nil {
		log.Println("Cannot connect DB: ", db_endpoint, err)
	} else {
		log.Println("Successfully connected DB: ", db_endpoint)
	}
	return
}
