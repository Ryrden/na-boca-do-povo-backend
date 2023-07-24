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
	FindAll() (int, dto.Response)
	FindById(ctx *gin.Context) (int, dto.Response)
	Create(ctx *gin.Context) (int, dto.Response)
	Update(ctx *gin.Context) (int, dto.Response)
	Delete(ctx *gin.Context) (int, dto.Response)
}

type controller struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() (int, dto.Response) {
	return c.service.FindAll()
}

func (c *controller) FindById(ctx *gin.Context) (int, dto.Response) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return http.StatusBadRequest, dto.InvalidIdResponse
	}
	return c.service.FindById(id)
}

func (c *controller) Create(ctx *gin.Context) (int, dto.Response) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		return http.StatusBadRequest, dto.Response{
			Success:  false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return c.service.Create(user)
}

func (c *controller) Update(ctx *gin.Context) (int, dto.Response) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return http.StatusBadRequest, dto.InvalidIdResponse
	}
	var newUser model.User
	if err := ctx.BindJSON(&newUser); err != nil {
		return http.StatusBadRequest, dto.Response{
			Success:  false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return c.service.Update(id, newUser)
}

func (c *controller) Delete(ctx *gin.Context) (int, dto.Response) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return http.StatusBadRequest, dto.InvalidIdResponse
	}
	return c.service.Delete(id)
}
