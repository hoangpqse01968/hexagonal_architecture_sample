package repository

import (
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/domain"
)

type UserRepository interface {
	GetUserById(id int64) (*domain.User, error)
}
