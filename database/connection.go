package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "evermos/config"
)

var db *gorm.DB
var err error

// ConnectDb connect to mysql
func Connect() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func GetDB() *gorm.DB {
	return db
}
