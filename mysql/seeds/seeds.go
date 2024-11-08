package seeds

import (
	"earn-expense/app/models"
	"github.com/jinzhu/gorm"
	"log"
)

func up(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}

func down(db *gorm.DB) {
	err := db.DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
}

func Run(db *gorm.DB) {
	down(db)
	up(db)

	for i, _ := range users {
		err := db.Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}
	}
}
