package carts

import (
	"evermos/repository"
)

func NewCartService(cartRepo repository.ICartRepository) *CartService {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func NewCartController(cartService ICartService) *CartController {
	return &CartController{
		cartService: cartService,
	}
}
