package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

func (c *DBConfig) NewDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Pass, c.Name, c.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	handleError(err)
	return db
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
