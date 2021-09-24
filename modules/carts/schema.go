package carts

import "evermos/models"

type CreateCartSchema struct {
	UserID   uint                 `form:"inv_id" json:"user_id" binding:"required" validate:"required,min=1"`
	CartItem models.CartInventory `form:"cart_item" json:"cart_item" binding:"required" validate:"required"`
}

type GeDataCartSchema struct {
	UserID uint `form:"inv_id" json:"user_id" binding:"required" validate:"required,min=1"`
}

type GeDataDetailCartSchema struct {
	UserID uint `form:"inv_id" json:"user_id" binding:"required" validate:"required,min=1"`
	CartID uint `form:"cart_id" json:"cart_id" binding:"required" validate:"required,min=1"`
}

type AddProductCartSchema struct {
	UserID   uint                 `form:"inv_id" json:"user_id" binding:"required" validate:"required,min=1"`
	CartID   uint                 `form:"cart_id" json:"cart_id" binding:"required" validate:"required,min=1"`
	CartItem models.CartInventory `form:"cart_item" json:"cart_item" binding:"required" validate:"required"`
}

type UpdateProductCartSchema struct {
	UserID   uint                 `form:"inv_id" json:"user_id" binding:"required" validate:"required,min=1"`
	CartID   uint                 `form:"cart_id" json:"cart_id" binding:"required" validate:"required,min=1"`
	CartItem models.CartInventory `form:"cart_item" json:"cart_item" binding:"required" validate:"required"`
}

type CheckoutCartSchema struct {
	UserID           uint `form:"inv_id" json:"user_id" binding:"required" validate:"required,min=1"`
	CartID           uint `form:"cart_id" json:"cart_id" binding:"required" validate:"required,min=1"`
	AmountExpedition uint `form:"amount_expedition" json:"amount_expedition" binding:"required" validate:"min=0"`
}
