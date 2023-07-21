package service

import (
	"github.com/ryrden/na-boca-do-povo-backend/model"
	"github.com/ryrden/na-boca-do-povo-backend/repository"
)

type UserService interface {
	FindAll() []model.User
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		repository: userRepository,
	}
}

func (service *userService) FindAll() []model.User {
	return service.repository.FindAll()
}
