package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"

	models "evermos/models"
)

type IAuthRepository interface {
	FindByPk(*models.User) error
	FindUserByEmail(email string) (res models.User, err error)
	Create(user models.User) (res models.User, err error)
	UpdateLastLogin(user models.User) (err error)
}

type AuthRepository struct {
	Conn *gorm.DB
}

func (m *AuthRepository) FindUserByEmail(email string) (res models.User, err error) {
	var user models.User
	result := m.Conn.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return models.User{}, errors.New("incorrect email")
	}
	return user, nil
}

func (m *AuthRepository) FindByPk(user *models.User) error {
	result := m.Conn.First(&user)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
func (m *AuthRepository) UpdateLastLogin(user models.User) (err error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	user.LastLoginAt = time.Now().In(loc)
	result := m.Conn.Model(user).Where("email = ?", user.Email).Updates(&user)
	if result.RowsAffected <= 0 {
		return errors.New("UpdateLastLogin error")
	}
	return nil
}

func (m *AuthRepository) Create(user models.User) (res models.User, err error) {
	result := m.Conn.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
