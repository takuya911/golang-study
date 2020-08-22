package usecase

import (
	"context"

	"github.com/takuya911/golang-study/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase func
func NewUserUsecase(u domain.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (domain.User, error) {
	result, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		panic(err)
	}
	return result, nil
}

func (u *userUsecase) Store(ctx context.Context, form *domain.User) (string, error) {
	if err := u.userRepo.Store(ctx, form); err != nil {
		panic(err)
	}
	return "成功", nil
}

func (u *userUsecase) Update(ctx context.Context, form *domain.User) (string, error) {
	if err := u.userRepo.Update(ctx, form); err != nil {
		panic(err)
	}
	return "成功", nil
}

// func (u *userUsecase) Delete(ctx context.Context, id int) (string, error) {
// 	if result := u.db.Where("id = ?", id).Delete(&domain.User{}); result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return "成功", nil

// }
