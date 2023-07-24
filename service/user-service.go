package service

import (
	"net/http"

	"github.com/ryrden/na-boca-do-povo-backend/dto"
	"github.com/ryrden/na-boca-do-povo-backend/model"
	"github.com/ryrden/na-boca-do-povo-backend/repository"
)

type UserService interface {
	FindAll() (int, dto.Response)
	FindById(id uint64) (int ,dto.Response)
	Create(newUser model.User) (int ,dto.Response)
	Update(id uint64, newUser model.User) (int ,dto.Response)
	Delete(id uint64) (int ,dto.Response)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		repository: userRepository,
	}
}

func (service *userService) FindAll() (int, dto.Response) {
	users, err := service.repository.FindAll()
	if err != nil {
		return http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return http.StatusOK, dto.Response{
		Success: true,
		Message: "users found successfully",
		Data:    users,
	}
}

func (service *userService) FindById(id uint64) (int, dto.Response) {
	user, err := service.repository.FindById(id)
	if err != nil {
		return http.StatusNotFound, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return http.StatusOK, dto.Response{
		Success: true,
		Message: "user found successfully",
		Data:    user,
	}
}

func (service *userService) Create(user model.User) (int, dto.Response) {
	newUser, err := service.repository.Create(user)
	if err != nil {
		return http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return http.StatusOK, dto.Response{
		Success: true,
		Message: "user created successfully",
		Data:    newUser,
	}
}

func (service *userService) Update(id uint64, newUser model.User) (int, dto.Response) {
	newUser, err := service.repository.Update(id, newUser)
	if err != nil {
		return http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return http.StatusOK, dto.Response{
		Success: true,
		Message: "user updated successfully",
		Data:    newUser,
	}
}

func (service *userService) Delete(id uint64) (int, dto.Response) {
	deletedUser, err := service.repository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return http.StatusOK, dto.Response{
		Success: true,
		Message: "user deleted successfully",
		Data:    deletedUser,
	}
}
