package model

import (
	"gorm.io/datatypes"
	"time"
)

type User struct {
	ID                      uint64                                      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                    string                                      `json:"name" gorm:"not null" binding:"required"`
	Email                   string                                      `json:"email" gorm:"not null" binding:"required,email"`
	Password                string                                      `json:"password" gorm:"not null" binding:"required"`
	CreatedAt               time.Time                                   `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt               time.Time                                   `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	FavoriteCongressPersons datatypes.JSONSlice[FavoriteCongressPerson] `json:"favorite_congress_persons" gorm:"type:jsonb"`
}
