package domain

import (
	"time"

	"gorm.io/gorm"
)



type User struct {
	gorm.Model
	Name string
	Email string
	createdAt time.Time `gorm:"autoCreateTime"`
	passwordHash string
}