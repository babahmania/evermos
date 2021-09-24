package models

import (
	"time"
)

// ProductOrder type
type ProductOrder struct {
	POID        uint      `gorm:"column:po_id; primary_key" json:"po_id" swaggerignore:"true"`
	InvID       uint      `gorm:"column:inv_id; primary_key" json:"inv_id" swaggerignore:"true"`
	SupplierID  uint      `gorm:"type:int; unsigned; not null; default:1;" json:"supplier_id" swaggerignore:"true"`
	OrderType   string    `json:"order_type"`
	OrderDate   time.Time `gorm:"timestamp; not null;" json:"order_date" swaggerignore:"true"`
	QtyOrder    uint      `gorm:"type:int; unsigned; not null; default:0;" json:"qty_order" swaggerignore:"true"`
	AmountPrice uint      `gorm:"type:int; unsigned; not null; default:0;" json:"amount_price" swaggerignore:"true"`
	AmountDisc  uint      `gorm:"type:int; unsigned; not null; default:0;" json:"amount_disc" swaggerignore:"true"`
	IsPromo     string    `gorm:"type:char(1); not null; default:'0'; index:idx_inv_promo" json:"is_promo" swaggerignore:"true"`
}
