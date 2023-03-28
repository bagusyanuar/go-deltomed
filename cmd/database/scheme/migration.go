package scheme

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Division{})
	database.AutoMigrate(&Location{})
	database.AutoMigrate(&Complaint{})
	log.Println("success migrate database")
}
