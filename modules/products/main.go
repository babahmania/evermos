package products

import (
	"evermos/repository"
)

func NewProductService(productRepo repository.IProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func NewProductController(productService IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
