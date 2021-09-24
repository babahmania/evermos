package products

import (
	"evermos/dto"
	"evermos/models"
	"evermos/repository"
)

type IProductService interface {
	ListProducts(pagi *dto.Pagination) ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	AddStockProduct(product models.Product) (models.Product, error)
}

type ProductService struct {
	productRepo repository.IProductRepository
}

func (a *ProductService) ListProducts(pagi *dto.Pagination) (res []models.Product, err error) {
	return a.productRepo.ListProducts(pagi)
}

func (a *ProductService) CreateProduct(product models.Product) (res models.Product, err error) {
	return a.productRepo.CreateProduct(product)
}

func (a *ProductService) AddStockProduct(product models.Product) (res models.Product, err error) {
	return a.productRepo.AddStockProduct(product)
}
