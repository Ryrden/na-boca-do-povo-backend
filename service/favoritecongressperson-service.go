package service

import (
	"encoding/json"
	"net/http"

	"github.com/ryrden/na-boca-do-povo-backend/dto"
	"github.com/ryrden/na-boca-do-povo-backend/repository"
)

type FavoriteCongressPersonService interface {
	AddFavoriteCongressPerson(id uint64, congressPerson json.RawMessage) (int, dto.Response)
}

type favoriteCongressPersonService struct {
	repository repository.FavoriteCongressPersonRepository
}

func NewFavoriteCongressPersonService(repository repository.FavoriteCongressPersonRepository) FavoriteCongressPersonService {
	return &favoriteCongressPersonService{
		repository: repository,
	}
}

func (service *favoriteCongressPersonService) AddFavoriteCongressPerson(id uint64, congressPerson json.RawMessage) (int, dto.Response) {
	newCongressPerson, err := service.repository.AddFavoriteCongressPerson(id, congressPerson)
	if err != nil {
		return http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	
	return http.StatusOK, dto.Response{
		Success: true,
		Message: "favorite congress person added successfully",
		Data:    newCongressPerson,
	}
}
