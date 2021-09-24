package models

import (
	"time"

	"gorm.io/gorm"
)

// Product type
type Product struct {
	InvID       uint           `gorm:"column:inv_id; primary_key" json:"inv_id" swaggerignore:"true"`
	SupplierID  uint           `gorm:"type:int; unsigned; not null; default:1;index:idx_inv_supplier_id" json:"supplier_id" swaggerignore:"true"`
	Name        string         `gorm:"type:varchar(255);unique;not null;unique_index" json:"name"`
	Description string         `gorm:"column:description; type:varchar(255);not null" json:"description"`
	QtyStock    uint           `gorm:"type:int; unsigned; not null; default:0; index:idx_inv_qty_stock" json:"qty_stock" swaggerignore:"true"`
	QtyMinStock uint           `gorm:"type:int; unsigned; not null; default:1;" json:"qty_min_stock" swaggerignore:"true"`
	QtyMinCart  uint           `gorm:"type:int; unsigned; not null; default:1;" json:"qty_min_cart" swaggerignore:"true"`
	QtyPavorit  uint           `gorm:"type:int; unsigned; not null; default:0;" json:"qty_pavorit" swaggerignore:"true"`
	AmountPrice uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_price" swaggerignore:"true"`
	AmountDisc  uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_disc" swaggerignore:"true"`
	QtyLike     uint           `gorm:"type:int; unsigned; not null; default:0;" json:"qty_like" swaggerignore:"true"`
	IsPromo     string         `gorm:"type:char(1); not null; default:'0'; index:idx_inv_promo" json:"is_promo" swaggerignore:"true"`
	IsActive    string         `gorm:"type:char(1); not null; default:'1'; index:idx_inv_status" json:"is_active" swaggerignore:"true"`
	CreatedAt   time.Time      `gorm:"timestamp; not null;" json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time      `gorm:"timestamp; not null;" json:"updated_at" swaggerignore:"true"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}
