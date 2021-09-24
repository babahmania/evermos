package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User type
type User struct {
	ID          uint           `gorm:"column:user_id; primary_key" json:"user_id" swaggerignore:"true"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name" swaggerignore:"true"`
	Email       string         `gorm:"type:varchar(100);unique;not null;unique_index" json:"email"`
	Password    string         `gorm:"type:varchar(255);not null" json:"password"`
	LastLoginAt time.Time      `gorm:"timestamp;" json:"last_login_at" swaggerignore:"true"`
	IsActive    string         `gorm:"type:char(1); not null; default:'1'; index:idx_user_status" json:"is_active" swaggerignore:"true"`
	CreatedAt   time.Time      `gorm:"timestamp; not null;" json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time      `gorm:"timestamp; not null;" json:"updated_at" swaggerignore:"true"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
	return
}
