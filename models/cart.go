package models

import (
	"time"

	"gorm.io/gorm"
)

// Cart type
type Cart struct {
	CartID      uint           `gorm:"column:cart_id; primary_key" json:"cart_id"`
	UserID      uint           `json:"user_id"`
	QtyItem     uint           `gorm:"type:int; unsigned; not null; default:0;" json:"qty_item" swaggerignore:"true"`
	AmountPrice uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_price" swaggerignore:"true"`
	AmountDisc  uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_disc" swaggerignore:"true"`
	AmountTotal uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_total" swaggerignore:"true"`
	IsActive    string         `gorm:"type:char(1); not null; default:'1'; index:idx_inv_status" json:"is_active" swaggerignore:"true"`
	CreatedAt   time.Time      `gorm:"timestamp; not null;" json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time      `gorm:"timestamp; not null;" json:"updated_at" swaggerignore:"true"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}

type CartInventory struct {
	SupplierID  uint   `json:"supplier_id" swaggerignore:"true" binding:"required" validate:"required,min=1"`
	InvID       uint   `json:"inv_id" swaggerignore:"true" binding:"required" validate:"required,min=1"`
	Name        string `json:"name" swaggerignore:"true" binding:"required" validate:"required,min=1"`
	QtyIem      uint   `json:"qty_item" swaggerignore:"true" binding:"required" validate:"required,min=1"`
	AmountPrice uint   `json:"amount_price" swaggerignore:"true" binding:"required" validate:"required,min=1"`
	AmountDisc  uint   `gorm:"column:amount_disc;" json:"amount_disc" swaggerignore:"true" binding:"required" validate:"min=0"`
}

type CartDetail struct {
	CartID      uint           `gorm:"column:cart_id;" json:"cart_id" swaggerignore:"true"`
	CartNo      uint           `gorm:"column:cart_no; type:int; unsigned; not null; default:1;" json:"cart_no" swaggerignore:"true"`
	SupplierID  uint           `gorm:"type:int; unsigned; not null; default:1;" json:"supplier_id" swaggerignore:"true"`
	InvID       uint           `gorm:"column:inv_id;" json:"inv_id" swaggerignore:"true"`
	Name        string         `gorm:"type:varchar(255);" json:"name"`
	QtyOrder    uint           `gorm:"column:qty_order;" json:"qty_order" swaggerignore:"true"`
	AmountPrice uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_price" swaggerignore:"true"`
	AmountDisc  uint           `gorm:"column:amount_disc; type:int; unsigned; not null; default:0;" json:"amount_disc" swaggerignore:"true"`
	AmountTotal uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_total" swaggerignore:"true"`
	Noted       string         `gorm:"type:varchar(255);" json:"noted"`
	IsPromo     string         `gorm:"type:char(1); not null; default:'0';" json:"is_promo" swaggerignore:"true"`
	IsActive    string         `gorm:"type:char(1); not null; default:'1';" json:"is_active" swaggerignore:"true"`
	CreatedAt   time.Time      `gorm:"timestamp; not null;" json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time      `gorm:"timestamp; not null;" json:"updated_at" swaggerignore:"true"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}

type CartDetailData struct {
	CartID     uint `gorm:"column:cart_id; primary_key" json:"cart_id" swaggerignore:"true"`
	UserID     uint `json:"user_id" swaggerignore:"true"`
	HeaderData Cart
	OrderData  CartDetail
}
