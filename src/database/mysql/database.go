package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:Qwerty123@tcp(127.0.0.1:5432)/mantenimiento"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	handleError(err)
	return db
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
