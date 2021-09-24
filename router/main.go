package router

import (
	_ "evermos/docs"
	authModule "evermos/modules/auth"
	cartModule "evermos/modules/carts"
	productModule "evermos/modules/products"

	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.Handler)
	authModule.RegisterRouters(app)
	productModule.RegisterRouters(app)
	cartModule.RegisterRouters(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
