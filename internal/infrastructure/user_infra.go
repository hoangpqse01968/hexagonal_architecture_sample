package infrastructure

import (
	"errors"

	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/domain"
)

type UserInfrastructure struct {
	users map[int64]*domain.User
}

func NewUserInfrastructure() *UserInfrastructure {
	users := make(map[int64]*domain.User, 3)
	users[1] = &domain.User{Id: 1, Name: "A", Email: "A@gmail.com", Address: "Test address 1"}
	users[2] = &domain.User{Id: 2, Name: "B", Email: "B@gmail.com", Address: "Test address 2"}
	users[3] = &domain.User{Id: 3, Name: "C", Email: "C@gmail.com", Address: "Test address 3"}
	return &UserInfrastructure{
		users: users,
	}
}

func (u *UserInfrastructure) GetUserById(id int64) (*domain.User, error) {
	if user, ok := u.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("not found")
}
