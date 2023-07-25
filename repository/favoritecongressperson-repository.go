package repository

import (
	"encoding/json"
	"fmt"

	"github.com/ryrden/na-boca-do-povo-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FavoriteCongressPersonRepository interface {
	FindAll() ([]model.FavoriteCongressPerson, error)
	AddFavoriteCongressPerson(id uint64, congressPerson json.RawMessage) (json.RawMessage, error)
}

type FCPRepositoryDatabase struct {
	connection *gorm.DB
}

func NewFavoriteCongressPersonRepository() FavoriteCongressPersonRepository {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Brazil/East"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return &FCPRepositoryDatabase{
		connection: db,
	}

func (db *FCPRepositoryDatabase) FindAll() ([]model.FavoriteCongressPerson, error) {
	//TODO: Implement
	return model.FavoriteCongressPerson{}, nil
}

func (db *FCPRepositoryDatabase) AddFavoriteCongressPerson(id uint64, congressPerson json.RawMessage) (json.RawMessage, error) {
	// var user model.User
	// dbResponse := db.connection.First(&user, id)
	// if dbResponse.Error != nil {
	// 	return json.RawMessage{}, fmt.Errorf("user with id '%d' not found", id)
	// }

	// // REVIEW: Is this the best way to do this?
	// user.FavoriteCongressPersons = append(user.FavoriteCongressPersons, congressPerson)

	// if err := db.connection.Save(&user).Error; err != nil {
	// 	return json.RawMessage{}, fmt.Errorf("failed to add favorite congress person to user with id '%d'", id)
	// }
	//TODO: Implement
	return json.RawMessage{}, nil
}
