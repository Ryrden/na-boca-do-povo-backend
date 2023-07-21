package model

type User struct {
	ID                     int    `json:"id" gorm:"primaryKey"`
	Name                   string `json:"name" gorm:"not null" binding:"required"`
	Email                  string `json:"email" gorm:"not null" binding:"required"`
	Password               string `json:"password" gorm:"not null" binding:"required"`
	CreatedAt              string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt              string `json:"updated_at" gorm:"autoUpdateTime"`
	FavoriteCongressperson string `json:"favorite_congressperson"` //this is a json string
}
