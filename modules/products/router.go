package products

import (
	"evermos/middlewares"
	"evermos/repository"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(app fiber.Router) {
	productRepo := repository.NewProductRepository()
	productService := NewProductService(productRepo)
	productController := NewProductController(productService)

	router := app.Group("/api/v1/products").Use(middlewares.AuthorizeJWT)
	router.Get("/", productController.listProducts)
	router.Post("/", productController.createProduct)
	router.Post("/add-stock", productController.addStockProduct)
	//router.Post("/check-stock-min", productController.createProduct)
}
