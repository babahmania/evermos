package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"evermos/models"

	"gorm.io/gorm"
)

// Migration for table products, models product
func SeedProduct(db *gorm.DB) {
	jsonFile, err := os.Open("database/data/products.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("opened products.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var products []models.Product
	json.Unmarshal([]byte(byteValue), &products)
	db.Create(&products)
}
