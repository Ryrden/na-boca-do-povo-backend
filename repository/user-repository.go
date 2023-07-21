package repository

import (
	"github.com/ryrden/na-boca-do-povo-backend/model"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type UserRepository interface {
	FindAll() []model.User
}

type database struct {
	connection *gorm.DB
}

// postGres
func NewUserRepository() UserRepository {
	url := "jdbc:postgresql://localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	return &database{
		connection: db,
	}
}

func (db *database) FindAll() []model.User {
	var users []model.User
	db.connection.Find(&users)
	return users
}
