package usecase

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/takuya911/golang-study/domain"
)

type userUsecase struct {
	db *gorm.DB
}

// NewUserUsecase func
func NewUserUsecase(db *gorm.DB) *userUsecase {
	return &userUsecase{db}
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	if result := u.db.Where("id = ?", id).Find(&user); result.Error != nil {
		panic(result.Error)
	}
	return user, nil
}

func (u *userUsecase) Store(ctx context.Context, form *domain.User) (string, error) {
	if result := u.db.Create(&form); result.Error != nil {
		panic(result.Error)
	}
	return "成功", nil
}

func (u *userUsecase) Update(ctx context.Context, form *domain.User) (string, error) {
	if result := u.db.Model(&form).Where("id = ?", form.ID).Updates(&form); result.Error != nil {
		panic(result.Error)
	}
	return "成功", nil
}

func (u *userUsecase) Delete(ctx context.Context, id int) (string, error) {
	if result := u.db.Where("id = ?", id).Delete(&domain.User{}); result.Error != nil {
		panic(result.Error)
	}
	return "成功", nil

}
