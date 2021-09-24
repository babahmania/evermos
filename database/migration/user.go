package migration

import (
	"evermos/models"

	"gorm.io/gorm"
)

// Migration for table users, models User
func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
