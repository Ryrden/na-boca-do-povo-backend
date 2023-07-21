package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryrden/na-boca-do-povo-backend/controller"
	"github.com/ryrden/na-boca-do-povo-backend/model"
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
	ctx.JSON(http.StatusOK, user)
}

//TODO: implement
func (api *UserApi) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, model.NotImplementedResponse)
}

//TODO: implement
func (api *UserApi) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, model.NotImplementedResponse)
}

//TODO: implement
func (api *UserApi) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, model.NotImplementedResponse)
}

//TODO: implement
func (api *UserApi) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, model.NotImplementedResponse)
}

