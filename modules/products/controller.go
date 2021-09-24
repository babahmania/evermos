package products

import (
	"evermos/dto"
	"evermos/middlewares"
	"evermos/models"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService IProductService
}

type IProductController interface {
	listProducts(ctx *fiber.Ctx) error
	createProduct(ctx *fiber.Ctx) error
	AddStockProduct(ctx *fiber.Ctx) error
}

// listProducts func gets all item products.
// @Description Get all item products.
// @Summary get all item products
// @Tags Inventory
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Security ApiKeyAuth
// @Router /api/v1/products [get]
func (a *ProductController) listProducts(c *fiber.Ctx) error {
	pagi := new(dto.Pagination)
	if err := c.QueryParser(pagi); err != nil {
		return err
	}
	pagi.FillDefault()
	books, _ := a.productService.ListProducts(pagi)
	pagi.Update()
	c.Links(
		"http://api.example.com/users?page=2", "next",
		"http://api.example.com/users?page=5", "last",
	)
	return c.Status(200).JSON(fiber.Map{
		"meta": pagi.GetMeta(),
		"data": books,
	})
}

// createProduct func create item.
// @Description create product item.
// @Summary create product item
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body CreateProductSchema true "Product Json"
// @Success 200 {object} models.Product
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Duplicate entry"
// @Security ApiKeyAuth
// @Router /api/v1/products [post]
func (a *ProductController) createProduct(c *fiber.Ctx) error {
	var form CreateProductSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}

	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}

	product := models.Product{Name: form.Name, Description: form.Description}
	newProduct, err := a.productService.CreateProduct(product)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"data": newProduct})
}

// addStockProduct func update stock produc item.
// @Description update stock produc item.
// @Summary update stock produc item
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body UpdateStockProductSchema true "Product Json"
// @Success 200 {object} models.Product
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Duplicate entry"
// @Security ApiKeyAuth
// @Router /api/v1/products/add-stock [post]
func (a *ProductController) addStockProduct(c *fiber.Ctx) error {
	var form UpdateStockProductSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}

	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}
	product := models.Product{InvID: form.InvID, SupplierID: form.SupplierID, QtyStock: form.QtyStock}
	newStock, err := a.productService.AddStockProduct(product)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"data": newStock})
}
