package auth

import (
	"evermos/middlewares"
	"evermos/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterRouters(app fiber.Router) {
	authRepo := repository.NewAuthRepository()
	authService := NewAuthService(authRepo)
	authController := NewAutController(authService)

	router := app.Group("/api/v1/users", logger.New())
	router.Post("/login", authController.login)
	router.Post("/register", authController.register)
	router.Get("/profile", middlewares.AuthorizeJWT, authController.getUser)
}
