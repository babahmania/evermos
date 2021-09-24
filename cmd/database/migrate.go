package main

import (
	"os"

	"evermos/database"
	"evermos/database/migration"
	"evermos/database/seeder"
)

func main() {
	err := database.Connect()
	if err != nil {
		panic("Failed to connect database")
	}
	db := database.GetDB()

	migration.MigrateUser(db)
	migration.MigrateProduct(db)

	if len(os.Args) <= 2 {
		return
	}
	option := os.Args[1]
	if option == "seed" {
		switch os.Args[2] {
		case "users":
			seeder.SeedUser(db)
		case "products":
			seeder.SeedProduct(db)
		}
	}
}
