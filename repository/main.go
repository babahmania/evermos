package repository

import "evermos/database"

func NewAuthRepository() IAuthRepository {
	return &AuthRepository{database.GetDB()}
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{database.GetDB()}
}

func NewCartRepository() ICartRepository {
	return &CartRepository{database.GetDB()}
}
