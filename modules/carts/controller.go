package carts

import (
	"evermos/dto"
	"evermos/middlewares"
	"evermos/models"

	"github.com/gofiber/fiber/v2"
)

type CartController struct {
	cartService ICartService
}

type ICartController interface {
	listCarts(ctx *fiber.Ctx) error
	getBadgeCart(ctx *fiber.Ctx) error
	getDetailCart(ctx *fiber.Ctx) error
	createCart(ctx *fiber.Ctx) error
	checkoutCart(ctx *fiber.Ctx) error
	addProductCart(ctx *fiber.Ctx) error
	updateProductCart(ctx *fiber.Ctx) error
}

// listCarts func gets all item carts.
// @Description Get all item carts.
// @Summary get all item carts
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} models.Cart
// @Security ApiKeyAuth
// @Router /api/v1/carts [get]
func (a *CartController) listCarts(c *fiber.Ctx) error {
	pagi := new(dto.Pagination)
	if err := c.QueryParser(pagi); err != nil {
		return err
	}
	pagi.FillDefault()
	books, _ := a.cartService.ListCarts(pagi)
	pagi.Update()
	c.Links(
		"http://158.140.191.182:50212/api/v1/carts?page=2", "next",
		"http://158.140.191.182:50212/api/v1/carts?page=5", "last",
	)
	return c.Status(200).JSON(fiber.Map{
		"meta": pagi.GetMeta(),
		"data": books,
	})
}

// getDetailCart func get value badge cart.
// @Description get value badge cart return value count of item quantity.
// @Summary get value badge cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body GeDataDetailCartSchema true "Cart Json"
// @Success 200 {object} models.Cart
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Cart is already open status"
// @Security ApiKeyAuth
// @Router /api/v1/carts/detail [get]
func (a *CartController) getDetailCart(c *fiber.Ctx) error {
	var form GeDataDetailCartSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}
	cart := models.CartDetailData{UserID: form.UserID, CartID: form.CartID}
	cartDetailData, err := a.cartService.GetDetailCart(cart)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"data": cartDetailData})
}

// getBadgeCart func get value badge cart. i'm sorry only run in postman
// @Description get value badge cart return value count of item quantity.
// @Summary get value badge cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body GeDataCartSchema false "Cart Json"
// @Success 200 {var} badgeCart
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Cart is already open status"
// @Security ApiKeyAuth
// @Router /api/v1/carts/badge [get]
func (a *CartController) getBadgeCart(c *fiber.Ctx) error {
	var form GeDataCartSchema
	var badgeCart int64 = 0
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}
	cart := models.Cart{UserID: form.UserID}
	badgeCart, err := a.cartService.GetBadgeCart(cart)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"data": badgeCart})
}

// createCart func create new cart.
// @Description create new cart header.
// @Summary create new cart header.
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body CreateCartSchema true "Cart Json"
// @Success 200 {object} models.Cart
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Cart is already open status"
// @Security ApiKeyAuth
// @Router /api/v1/carts [post]
func (a *CartController) createCart(c *fiber.Ctx) error {
	var form CreateCartSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}
	//var qtyItem uint = 1
	amountPrice := form.CartItem.QtyIem * form.CartItem.AmountPrice
	amountDisc := form.CartItem.QtyIem * form.CartItem.AmountDisc
	amountTotal := amountPrice - amountDisc

	cart := models.Cart{UserID: form.UserID, QtyItem: form.CartItem.QtyIem, AmountPrice: amountPrice, AmountDisc: amountDisc, AmountTotal: amountTotal}
	cartItem := models.CartDetail{SupplierID: form.CartItem.SupplierID, InvID: form.CartItem.InvID, Name: form.CartItem.Name,
		QtyOrder: form.CartItem.QtyIem, AmountPrice: form.CartItem.AmountPrice, AmountDisc: form.CartItem.AmountDisc, AmountTotal: amountTotal}

	newCart, err := a.cartService.CreateCart(cart, cartItem)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"data": newCart})
}

// addProductCart func add new product to cart.
// @Description add new product to cart.
// @Summary add new product to cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body AddProductCartSchema true "Cart Json"
// @Success 200 {object} models.Cart
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Duplicate entry"
// @Security ApiKeyAuth
// @Router /api/v1/carts/add-product [post]
func (a *CartController) addProductCart(c *fiber.Ctx) error {
	var form AddProductCartSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}
	amountPrice := form.CartItem.QtyIem * form.CartItem.AmountPrice
	amountDisc := form.CartItem.QtyIem * form.CartItem.AmountDisc
	amountTotal := amountPrice - amountDisc

	cart := models.Cart{CartID: form.CartID, UserID: form.UserID, QtyItem: form.CartItem.QtyIem, AmountPrice: amountPrice, AmountDisc: amountDisc, AmountTotal: amountTotal}

	//cartItem := new(models.CartDetail)
	//cartItem.InvID = form.CartItem.InvID
	cartItem := models.CartDetail{CartID: form.CartID, SupplierID: form.CartItem.SupplierID, InvID: form.CartItem.InvID, Name: form.CartItem.Name,
		QtyOrder: form.CartItem.QtyIem, AmountPrice: form.CartItem.AmountPrice, AmountDisc: form.CartItem.AmountDisc, AmountTotal: amountTotal}
	newCart, err := a.cartService.AddProductCart(cart, cartItem)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"data": newCart})
}

// updateProductCart func update qty order product item in cart.
// @Description update qty order product item in cart.
// @Summary update qty order product item in cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body UpdateProductCartSchema true "Cart Json"
// @Success 200 {object} models.Cart
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Duplicate entry"
// @Security ApiKeyAuth
// @Router /api/v1/carts/update-product [post]
func (a *CartController) updateProductCart(c *fiber.Ctx) error {
	var form UpdateProductCartSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	return nil
}

// checkoutCart func create new checkout cart by id.
// @Description create new checkout cart by id.
// @Summary create new checkout cart by id.
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body CheckoutCartSchema true "Cart Json"
// @Success 200 {object} models.Cart
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Cart is already open status"
// @Security ApiKeyAuth
// @Router /api/v1/carts [post]
func (a *CartController) checkoutCart(c *fiber.Ctx) error {
	var form CheckoutCartSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}

	cart := models.Sales{UserID: form.UserID, SalesID: form.CartID, AmountExpedition: form.AmountExpedition}

	newCheckout, err := a.cartService.CheckoutCart(cart)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"data": newCheckout})

}
