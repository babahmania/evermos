package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	dto "evermos/dto"
	models "evermos/models"
)

type IProductRepository interface {
	ListProducts(pagi *dto.Pagination) (res []models.Product, err error)
	CreateProduct(product models.Product) (models.Product, error)
	AddStockProduct(product models.Product) (models.Product, error)
}

type ProductRepository struct {
	Conn *gorm.DB
}

func (m *ProductRepository) ListProducts(pagi *dto.Pagination) (res []models.Product, err error) {
	var products []models.Product
	var count int64
	m.Conn.Limit(pagi.Limit).Offset(pagi.Offset).Find(&products)
	m.Conn.Model(&models.Product{}).Count(&count)
	pagi.Total = count

	return products, nil
}

func (m *ProductRepository) CreateProduct(product models.Product) (res models.Product, err error) {
	result := m.Conn.Create(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (m *ProductRepository) AddStockProduct(product models.Product) (res models.Product, err error) {
	result := m.Conn.Model(product).Where("inv_id = ?", product.InvID).UpdateColumn("qty_stock", gorm.Expr("qty_stock + ?", product.QtyStock))
	if result.RowsAffected <= 0 {
		fmt.Println(result.Error)
		return models.Product{}, result.Error
	}
	var productStock models.Product
	newStock := m.Conn.Where("inv_id = ?", product.InvID).First(&productStock)
	if newStock.RowsAffected == 0 {
		return models.Product{}, errors.New("product not found")
	}
	return productStock, nil
}
