package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryrden/na-boca-do-povo-backend/dto"
	"github.com/ryrden/na-boca-do-povo-backend/model"
	"github.com/ryrden/na-boca-do-povo-backend/service"
)

type UserController interface {
	FindAll() []model.User
	FindById(ctx *gin.Context) dto.Response
	Create(ctx *gin.Context) dto.Response
	Update(ctx *gin.Context) dto.Response
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

func (c *controller) FindById(ctx *gin.Context) dto.Response {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return dto.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid id",
			Data:  nil,
		}
	}
	return c.service.FindById(id)
}

func (c *controller) Create(ctx *gin.Context) dto.Response {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		return dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:  nil,
		}
	}
	return c.service.Create(user)
}

func (c *controller) Update(ctx *gin.Context) dto.Response {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return dto.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid id",
			Data:  nil,
		}
	}
	var newUser model.User
	if err := ctx.BindJSON(&newUser); err != nil {
		return dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:  nil,
		}
	}
	return c.service.Update(id, newUser)
}