package models

import (
	"time"

	"gorm.io/gorm"
)

// Sales type
type Sales struct {
	SalesID          uint           `gorm:"column:sales_id; primary_key" json:"sales_id" swaggerignore:"true"`
	UserID           uint           `json:"user_id" swaggerignore:"true"`
	SalesDate        time.Time      `gorm:"timestamp; not null;" json:"sales_date" swaggerignore:"true"`
	InvoiceNo        string         `gorm:"column:sales_inv_no;" json:"sales_inv_no" swaggerignore:"true"`
	UserAddress      string         `gorm:"column:user_address;" json:"user_address" swaggerignore:"true"`
	StatusOrder      string         `json:"status_order" swaggerignore:"true"`
	QtyItem          uint           `gorm:"type:int; unsigned; not null; default:0;" json:"qty_item" swaggerignore:"true"`
	AmountPrice      uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_price" swaggerignore:"true"`
	AmountDisc       uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_disc" swaggerignore:"true"`
	AmountExpedition uint           `gorm:"column:amount_expedition; type:int; unsigned; not null; default:0;" json:"amount_expedition" swaggerignore:"true"`
	AmountTotal      uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_total" swaggerignore:"true"`
	IsActive         string         `gorm:"type:char(1); not null; default:'1'; index:idx_inv_status" json:"is_active" swaggerignore:"true"`
	CreatedAt        time.Time      `gorm:"timestamp; not null;" json:"created_at" swaggerignore:"true"`
	UpdatedAt        time.Time      `gorm:"timestamp; not null;" json:"updated_at" swaggerignore:"true"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}

type SalesDetail struct {
	SalesID     uint           `gorm:"column:sales_id;" json:"sales_id" swaggerignore:"true"`
	SalesNo     uint           `gorm:"column:cart_no; type:int; unsigned; not null; default:1;" json:"sales_no" swaggerignore:"true"`
	SupplierID  uint           `gorm:"type:int; unsigned; not null; default:1;" json:"supplier_id" swaggerignore:"true"`
	InvID       uint           `gorm:"column:inv_id;" json:"inv_id" swaggerignore:"true"`
	Name        string         `gorm:"type:varchar(255);" json:"name"`
	QtyOrder    uint           `gorm:"column:qty_order;" json:"qty_order" swaggerignore:"true"`
	AmountPrice uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_price" swaggerignore:"true"`
	AmountDisc  uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_disc" swaggerignore:"true"`
	AmountTotal uint           `gorm:"type:int; unsigned; not null; default:0;" json:"amount_total" swaggerignore:"true"`
	Noted       string         `gorm:"type:varchar(255);" json:"noted"`
	IsPromo     string         `gorm:"type:char(1); not null; default:'0';" json:"is_promo" swaggerignore:"true"`
	IsActive    string         `gorm:"type:char(1); not null; default:'1';" json:"is_active" swaggerignore:"true"`
	CreatedAt   time.Time      `gorm:"timestamp; not null;" json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time      `gorm:"timestamp; not null;" json:"updated_at" swaggerignore:"true"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}

type SalesDetailData struct {
	SalesID    uint `gorm:"column:sales_id; primary_key" json:"sales_id" swaggerignore:"true"`
	UserID     uint `json:"user_id" swaggerignore:"true"`
	HeaderData Sales
	OrderData  SalesDetail
}
