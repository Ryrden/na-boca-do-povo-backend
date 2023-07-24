package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryrden/na-boca-do-povo-backend/controller"
	"github.com/ryrden/na-boca-do-povo-backend/dto"
)

type FavoriteCongressPersonApi struct {
	favoriteCongressPersonController controller.FavoriteCongressPersonController
}

func NewFavoriteCongressPersonApi(favoriteCongressPersonController controller.FavoriteCongressPersonController) FavoriteCongressPersonApi {
	return FavoriteCongressPersonApi{
		favoriteCongressPersonController: favoriteCongressPersonController,
	}
}

func (api *FavoriteCongressPersonApi) GetFavoriteCongressPersons(ctx *gin.Context) {
	//TODO: implement
	ctx.JSON(http.StatusNotImplemented, dto.NotImplementedResponse)
}

func (api *FavoriteCongressPersonApi) AddFavoriteCongressPerson(ctx *gin.Context) {
	status, response := api.favoriteCongressPersonController.AddFavoriteCongressPerson(ctx)
	if status != http.StatusOK {
		ctx.JSON(status, response)
		return
	}
	ctx.JSON(status, response)
}

func (api *FavoriteCongressPersonApi) RemoveFavoriteCongressPerson(ctx *gin.Context) {
	//TODO: implement
	ctx.JSON(http.StatusNotImplemented, dto.NotImplementedResponse)
}