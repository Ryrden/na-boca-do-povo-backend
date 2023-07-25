package model

import (
	"time"
	"gorm.io/datatypes"
)

type User struct {
	ID                      uint64            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                    string            `json:"name" gorm:"not null" binding:"required"`
	Email                   string            `json:"email" gorm:"not null" binding:"required,email"`
	Password                string            `json:"password" gorm:"not null" binding:"required"`
	CreatedAt               time.Time         `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt               time.Time         `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	FavoriteCongressPersons datatypes.JSON    `json:"favorite_congresspersons" gorm:"type:jsonb;default:'[]'"`
}
