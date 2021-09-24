package carts

import (
	"evermos/dto"
	"evermos/models"
	"evermos/repository"
)

type ICartService interface {
	ListCarts(pagi *dto.Pagination) ([]models.Cart, error)
	GetBadgeCart(cart models.Cart) (int64, error)
	GetDetailCart(cart models.CartDetailData) (models.CartDetailData, error)
	CreateCart(cart models.Cart, cartDetail models.CartDetail) (models.Cart, error)
	AddProductCart(cart models.Cart, cartDetail models.CartDetail) (models.Cart, error)
	UpdateProductCart(cart models.Cart, cartDetail models.CartDetail) (models.Cart, error)
	CheckoutCart(cart models.Sales) (models.Sales, error)
}

type CartService struct {
	cartRepo repository.ICartRepository
}

func (a *CartService) ListCarts(pagi *dto.Pagination) (res []models.Cart, err error) {
	return a.cartRepo.ListCarts(pagi)
}

func (a *CartService) GetBadgeCart(cart models.Cart) (badgeCart int64, err error) {
	return a.cartRepo.GetBadgeCart(cart)
}

func (a *CartService) GetDetailCart(cart models.CartDetailData) (res models.CartDetailData, err error) {
	return a.cartRepo.GetDetailCart(cart)
}

func (a *CartService) CreateCart(cart models.Cart, cartDetail models.CartDetail) (res models.Cart, err error) {
	return a.cartRepo.CreateCart(cart, cartDetail)
}

func (a *CartService) AddProductCart(cart models.Cart, cartDetail models.CartDetail) (res models.Cart, err error) {
	return a.cartRepo.AddProductCart(cart, cartDetail)
}

func (a *CartService) UpdateProductCart(cart models.Cart, cartDetail models.CartDetail) (res models.Cart, err error) {
	return a.cartRepo.UpdateProductCart(cart, cartDetail)
}

func (a *CartService) CheckoutCart(cart models.Sales) (res models.Sales, err error) {
	return a.cartRepo.CheckoutCart(cart)
}
