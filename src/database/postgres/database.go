package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=root password=Qwerty123 dbname=mantenimiento port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	handleError(err)
	return db
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
