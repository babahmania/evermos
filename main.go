package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"evermos/config"
	"evermos/database"
	"evermos/middlewares"
	"evermos/router"
	"evermos/schedulers"
	"evermos/utils/queue"
	"evermos/utils/redis"
)

// @title evermos-online-store
// @version 1.0
// @description GO REST API MYSQL FIBER GORM
// @contact.name babahmania
// @contact.url https://github.com/babahmania/evermos
// @contact.email babahmania@gmail.com
// @host localhost:50212
// @securityDefinitions.apikey ApiKeyAuth
// @in input value format 'Bearer '+access_token
// @name Authorization
// @BasePath /

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	err = database.Connect()
	if err != nil {
		panic("Failed to connect database")
	}
	app := fiber.New(middlewares.HandleApiError())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	router.Init(app)
	schedulers.Init()
	redis.Init()
	queue.Init()

	log.Info("Starting API server at port " + config.PORT)
	app.Listen(":" + config.PORT)
}
