package carts

import (
	"evermos/middlewares"
	"evermos/repository"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(app fiber.Router) {
	cartRepo := repository.NewCartRepository()
	cartService := NewCartService(cartRepo)
	cartController := NewCartController(cartService)

	router := app.Group("/api/v1/carts").Use(middlewares.AuthorizeJWT)
	router.Get("/", cartController.listCarts)
	router.Get("/badge", cartController.getBadgeCart)
	router.Get("/detail", cartController.getDetailCart)
	router.Post("/", cartController.createCart)
	router.Post("/add-product", cartController.addProductCart)
	router.Post("/update-product", cartController.updateProductCart)
	router.Post("/checkout", cartController.checkoutCart)

	//when open page cart
	//router.Post("/check-stock-cart", cartController.checkStockCart)
}
