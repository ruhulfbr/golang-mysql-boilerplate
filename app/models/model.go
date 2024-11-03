package models

import "github.com/jinzhu/gorm"

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})

	return db
}
