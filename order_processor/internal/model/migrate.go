package model

import "gorm.io/gorm"

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&Order{})
}
