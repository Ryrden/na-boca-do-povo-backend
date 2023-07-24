package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryrden/na-boca-do-povo-backend/controller"
)

type UserApi struct {
	userController controller.UserController
}

func NewUserAPI(userController controller.UserController) *UserApi {
	return &UserApi{
		userController: userController,
	}
}

func (api *UserApi) GetUsers(ctx *gin.Context) {
	status, response := api.userController.FindAll()
	if status != http.StatusOK {
		ctx.JSON(status, response)
		return
	}
	ctx.JSON(status, response)
}

func (api *UserApi) GetUser(ctx *gin.Context) {
	status, response := api.userController.FindById(ctx)
	if status != http.StatusOK {
		ctx.JSON(status, response)
		return
	}
	ctx.JSON(status, response)
}

func (api *UserApi) Create(ctx *gin.Context) {
	status, response := api.userController.Create(ctx)
	if status != http.StatusOK {
		ctx.JSON(status, response)
		return
	}
	ctx.JSON(status, response)
}

func (api *UserApi) Update(ctx *gin.Context) {
	status, response := api.userController.Update(ctx)
	if status != http.StatusOK {
		ctx.JSON(status, response)
		return
	}
	ctx.JSON(status, response)
}

// TODO: implement
func (api *UserApi) Delete(ctx *gin.Context) {
	status, response := api.userController.Delete(ctx)
	if status != http.StatusOK {
		ctx.JSON(status, response)
		return
	}
	ctx.JSON(status, response)
}
