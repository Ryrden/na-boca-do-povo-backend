package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryrden/na-boca-do-povo-backend/dto"
	"github.com/ryrden/na-boca-do-povo-backend/model"
	"github.com/ryrden/na-boca-do-povo-backend/service"
)

type FavoriteCongressPersonController interface {
	FindAll(ctx *gin.Context) (int, dto.Response)
	AddFavoriteCongressPerson(ctx *gin.Context) (int, dto.Response)
	//RemoveFavoriteCongressPerson(ctx *gin.Context) dto.Response
	//
}

type favoriteCongressPersonController struct {
	service service.FavoriteCongressPersonService
}

func NewFavoriteCongressPersonController(service service.FavoriteCongressPersonService) FavoriteCongressPersonController {
	return &favoriteCongressPersonController{
		service: service,
	}
}

func (c *favoriteCongressPersonController) FindAll(ctx *gin.Context) (int, dto.Response) {
	return c.service.FindAll()
}

func (c *favoriteCongressPersonController) AddFavoriteCongressPerson(ctx *gin.Context) (int, dto.Response) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return http.StatusBadRequest, dto.InvalidIdResponse
	}
	var favoriteCongressPerson model.FavoriteCongressPerson
	if err := ctx.BindJSON(&favoriteCongressPerson); err != nil {
		return http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	favoriteCongressPersonJson, err := json.Marshal(favoriteCongressPerson)
	if err != nil {
		return http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	favoriteCongressPersonJson = json.RawMessage(favoriteCongressPersonJson)
	return c.service.AddFavoriteCongressPerson(id, favoriteCongressPersonJson)
}
