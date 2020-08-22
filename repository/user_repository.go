package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/takuya911/golang-study/domain"
)

type userRepo struct {
	Conn *gorm.DB
}

// NewUserRepository function
func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &userRepo{Conn}
}

// GetByUserID function
func (u *userRepo) GetByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	if result := u.Conn.Where("id = ?", id).Find(&user); result.Error != nil {
		panic(result.Error)
	}
	return user, nil
}
