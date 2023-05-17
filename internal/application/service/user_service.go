package service

import (
	"fmt"

	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/domain"
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/repository"
)

type UserService interface {
	GetUser(userid int64) (*domain.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(
	userRepo repository.UserRepository,
) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (u *UserServiceImpl) GetUser(userid int64) (*domain.User, error) {
	user, err := u.userRepo.GetUserById(userid)
	if err != nil {
		return nil, fmt.Errorf("failed to get user, err[%v]", err)
	}

	return user, nil
}
