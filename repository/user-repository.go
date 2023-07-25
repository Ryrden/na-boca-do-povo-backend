package repository

import (
	"fmt"

	"github.com/ryrden/na-boca-do-povo-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(id uint64) (model.User, error)
	Create(newUser model.User) (model.User, error)
	Update(id uint64, newUser model.User) (model.User, error)
	Delete(id uint64) (model.User, error)
	AddFavoriteCongressPerson(id uint64, congressPerson model.FavoriteCongressPerson) (model.User, error)
}

type UserRepositoryDatabase struct {
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
	return &UserRepositoryDatabase{
		connection: db,
	}
}

func (db *UserRepositoryDatabase) FindAll() ([]model.User, error) {
	var users []model.User
	dbResponse := db.connection.Order("id").Find(&users)
	if dbResponse.Error != nil {
		// REVIEW: is this the best way to handle this error?
		return nil, fmt.Errorf("error finding users: %s", dbResponse.Error)
	}
	return users, nil
}

func (db *UserRepositoryDatabase) FindById(id uint64) (model.User, error) {
	var user model.User
	dbResponse := db.connection.First(&user, id)
	if dbResponse.Error != nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("user with id '%d' not found", id)
	}
	return user, nil
}

func (db *UserRepositoryDatabase) Create(user model.User) (model.User, error) {
	var existingUser model.User
	dbResponse := db.connection.Where("email = ?", user.Email).First(&existingUser)
	if dbResponse.Error == nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("user with email '%s' already exists", user.Email)
	}

	dbResponse = db.connection.Create(&user)
	if dbResponse.Error != nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("error creating user: %s", dbResponse.Error)
	}
	return user, nil
}

func (db *UserRepositoryDatabase) Update(id uint64, newUser model.User) (model.User, error) {
	var updatedUser model.User
	dbResponse := db.connection.First(&updatedUser, id)
	if dbResponse.Error != nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("user with id '%d' not found", id)
	}

	var existingUser model.User
	dbResponse = db.connection.Where("email = ?", newUser.Email).First(&existingUser)
	if dbResponse.Error == nil && existingUser.ID != id {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("user with email '%s' already exists", newUser.Email)
	}

	//TODO: Use Mutex to threat concurrency
	updatedUser.Name = newUser.Name
	updatedUser.Email = newUser.Email
	updatedUser.Password = newUser.Password

	dbResponse = db.connection.Save(&updatedUser)
	if dbResponse.Error != nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("error updating user with id '%d: %s'", id, dbResponse.Error)
	}
	return updatedUser, nil
}

func (db *UserRepositoryDatabase) Delete(id uint64) (model.User, error) {
	var deletedUser model.User
	dbResponse := db.connection.First(&deletedUser, id)
	if dbResponse.Error != nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("user with id '%d' not found", id)
	}
	//TODO: Use Mutex to threat concurrency
	dbResponse = db.connection.Delete(&deletedUser)
	if dbResponse.Error != nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("error deleting user with id '%d': %s", id, dbResponse.Error)
	}
	return deletedUser, nil
}

func (db *UserRepositoryDatabase) AddFavoriteCongressPerson(id uint64, congressPerson model.FavoriteCongressPerson) (model.User, error) {
	var user model.User
	dbResponse := db.connection.First(&user, id)
	if dbResponse.Error != nil {
		// TODO: Create Error Handler 
		return model.User{}, fmt.Errorf("user with id '%d' not found", id)
	}
	var favoriteCongressPersons datatypes.JSONSlice[model.FavoriteCongressPerson]
	favoriteCongressPersons = user.FavoriteCongressPersons
	favoriteCongressPersons = append(favoriteCongressPersons, congressPerson)

	dbResponse = db.connection.Model(&user).Update("favorite_congress_persons", favoriteCongressPersons)
	if dbResponse.Error != nil {
		return model.User{}, fmt.Errorf("error adding favorite congress person to user with id '%d': %s", id, dbResponse.Error)
	}
	return user, nil
}
