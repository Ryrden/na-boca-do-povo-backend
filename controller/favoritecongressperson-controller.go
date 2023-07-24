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
	//GetFavoriteCongressPersons(ctx *gin.Context) dto.Response
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

//func (c *favoriteCongressPersonController) GetFavoriteCongressPersons(ctx *gin.Context) dto.Response {

func (c *favoriteCongressPersonController) AddFavoriteCongressPerson(ctx *gin.Context) (int, dto.Response) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return http.StatusBadRequest, dto.InvalidIdResponse
	}
	var congressPerson model.CongressPerson
	if err := ctx.BindJSON(&congressPerson); err != nil {
		return http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	//convert congressPerson to json.RawMessage
	congressPersonJson, err := json.Marshal(congressPerson)
	if err != nil {
		return http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	congressPersonJson = json.RawMessage(congressPersonJson)
	return c.service.AddFavoriteCongressPerson(id, congressPersonJson)
}
