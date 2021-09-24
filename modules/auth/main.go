package auth

import (
	"evermos/repository"
)

func NewAuthService(authRepo repository.IAuthRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

func NewAutController(authService IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}
