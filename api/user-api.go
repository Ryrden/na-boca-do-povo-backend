package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryrden/na-boca-do-povo-backend/controller"
	"github.com/ryrden/na-boca-do-povo-backend/dto"
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
	user := api.userController.FindAll()
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Users found successfully",
		Data:   user,
	})
}

// TODO: implement
func (api *UserApi) GetUser(ctx *gin.Context) {
	response := api.userController.FindById(ctx)
	if response.Status != http.StatusOK {
		ctx.JSON(response.Status, response)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// TODO: implement
func (api *UserApi) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, dto.NotImplementedResponse)
}

// TODO: implement
func (api *UserApi) Update(ctx *gin.Context) {
	response := api.userController.Update(ctx)
	if response.Status != http.StatusOK {
		ctx.JSON(response.Status, response)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// TODO: implement
func (api *UserApi) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, dto.NotImplementedResponse)
}
