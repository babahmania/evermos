package products

type CreateProductSchema struct {
	Name        string `form:"name" json:"name" binding:"required" validate:"required,min=4,max=32"`
	Description string `form:"description" json:"description" binding:"required" validate:"required,min=4"`
}

type UpdateStockProductSchema struct {
	InvID      uint `form:"inv_id" json:"inv_id" binding:"required" validate:"required,min=1"`
	SupplierID uint `form:"supplier_id" json:"supplier_id" binding:"required" validate:"required,min=1"`
	QtyStock   uint `form:"qty_stock" json:"qty_stock" binding:"required" validate:"required,min=1"`
}
