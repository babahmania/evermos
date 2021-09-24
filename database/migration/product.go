package migration

import (
	"evermos/models"

	"gorm.io/gorm"
)

// Migration for table products, models Product
func MigrateProduct(db *gorm.DB) {
	db.AutoMigrate(&models.Product{})
}
