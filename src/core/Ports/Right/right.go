package right

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func GetDB() *gorm.DB {
	if database == nil {
		log.Fatal(color.RedString("Error: Database Provider Null"))
	}
	return database
}

func InjectDB(db *gorm.DB) {
	if database == nil {
		database = db
		return
	}
	fmt.Println(color.YellowString("WARN: Database has been already provided!"))
}
