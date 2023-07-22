package repository

import (
	"fmt"

	"github.com/ryrden/na-boca-do-povo-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []model.User
	FindById(id uint64) (model.User, error)
	Create(newUser model.User) (model.User, error)
	Update(id uint64, newUser model.User) (model.User, error)
}

type database struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	//TODO: use environment variables
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Brazil/East"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	//TODO: use pagination
	db.connection.Order("id").Find(&users)
	return users
}

func (db *database) FindById(id uint64) (model.User, error) {
	var user model.User
	dbResponse := db.connection.First(&user, id)
	if dbResponse.Error != nil {
		// REVIEW: is this the best way to handle this error?
		return model.User{}, fmt.Errorf("user with id '%d' not found", id)
	}
	return user, nil
}

func (db *database) Create(user model.User) (model.User, error) {
	var existingUser model.User
	dbResponse := db.connection.Where("email = ?", user.Email).First(&existingUser)
	if dbResponse.Error == nil {
		// REVIEW: is this the best way to handle this error?
		return model.User{}, fmt.Errorf("user with email '%s' already exists", user.Email) 
	}
	
	dbResponse = db.connection.Create(&user)
	if dbResponse.Error != nil {
		// REVIEW: is this the best way to handle this error?
		return model.User{}, fmt.Errorf("error creating user: %s", dbResponse.Error)
	}
	return user, nil
}


func (db *database) Update(id uint64, newUser model.User) (model.User, error) {
	var updatedUser model.User
	dbResponse := db.connection.First(&updatedUser, id)
	if dbResponse.Error != nil {
		// REVIEW: is this the best way to handle this error?
		return model.User{}, fmt.Errorf("user with id '%d' not found", id)
	}
	
	//TODO: Use Mutex to threat concurrency
	updatedUser.Name = newUser.Name
	updatedUser.Email = newUser.Email
	updatedUser.Password = newUser.Password

	dbResponse = db.connection.Save(&updatedUser)
	if dbResponse.Error != nil {
		// REVIEW: is this the best way to handle this error?
		return model.User{}, fmt.Errorf("error updating user with id '%d: %s'", id, dbResponse.Error)
	}
	return updatedUser, nil
}
