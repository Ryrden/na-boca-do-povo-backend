package controller

import (
	model "github.com/ryrden/na-boca-do-povo-backend/model"
	"github.com/ryrden/na-boca-do-povo-backend/service"
)

type UserController interface {
	FindAll() []model.User
}

type controller struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.User {
	return c.service.FindAll()
}
