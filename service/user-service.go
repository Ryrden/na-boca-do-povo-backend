package service

import (
	"net/http"

	"github.com/ryrden/na-boca-do-povo-backend/dto"
	"github.com/ryrden/na-boca-do-povo-backend/model"
	"github.com/ryrden/na-boca-do-povo-backend/repository"
)

type UserService interface {
	FindAll() []model.User
	FindById(id uint64) dto.Response
	Update(id uint64, newUser model.User) dto.Response
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

func (service *userService) FindById(id uint64) dto.Response {
	user, err := service.repository.FindById(id)
	if err != nil {
		return dto.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:   nil,
		}
	}
	return dto.Response{
		Status:  http.StatusOK,
		Message: "user found successfully",
		Data:    user,
	}
}

func (service *userService) Update(id uint64, newUser model.User) dto.Response {
	newUser, err := service.repository.Update(id, newUser)
	if err != nil {
		return dto.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:   nil,
		}
	}
	return dto.Response{
		Status:  http.StatusOK,
		Message: "user updated successfully",
		Data:    newUser,
	}
}
