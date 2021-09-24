package auth

import (
	"errors"
	"strconv"

	"evermos/constants"
	"evermos/models"
	"evermos/repository"
	Jwt "evermos/utils/jwt"
	"evermos/utils/redis"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(email string, password string) (token string, err error)
	CreateUser(user models.User) (usr models.User, err error)
	GetUser(user *models.User) error
}

type AuthService struct {
	authRepo repository.IAuthRepository
}

// Login method to auth user and return access and refresh tokens.
// @Description Auth user and return access and refresh token.<br>Example value : {"email": "babahmania@gmail.com", "password": "admin123"}
// @Summary auth user and return access and refresh token
// @Tags User
// @Accept json
// @Produce json
// @Param input body models.User true "User Json"
// @Success 200 {string} status "ok"
// @Router /api/v1/users/login [post]
func (a *AuthService) Login(email string, password string) (token string, err error) {
	user, err := a.authRepo.FindUserByEmail(email)
	if err != nil || user.IsActive == constants.STATUS.INACTIVE {
		return "", errors.New("invalid email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}
	//update LastLoginAt
	err = a.authRepo.UpdateLastLogin(user)
	if err != nil {
		return "", errors.New("updating LastLoginAt error")
	}

	token = Jwt.GenerateAccessToken(user)
	redis.SaveToken(strconv.Itoa(int(user.ID)), token)
	return token, nil
}

// CreateUser func create user.
// @Description Create new user register.<br>Example value : {"name": "new user", "email": "username@gmail.com", "password": "password123"}
// @Summary create new user
// @Tags User
// @Accept json
// @Produce json
// @Param input body models.User true "User Json"
// @Success 200 {object} models.User
// @Failure 422 {string} string "Unprocessable Entity"
// @Failure 500 {string} string "Internal Server Error or Duplicate entry"
// @Router /api/v1/users/register [post]
func (a *AuthService) CreateUser(user models.User) (usr models.User, err error) {
	return a.authRepo.Create(user)
}

// GetUser func gets user by given detail user profile.
// @Description Get user by given detail user profile.
// @Summary get detail user profile
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /api/v1/users/profile [get]
func (a *AuthService) GetUser(user *models.User) error {
	return a.authRepo.FindByPk(user)
}
