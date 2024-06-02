package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uint    `gorm:"primaryKey;required"`
	Username  string  `validate:"required"`
	Email     string  `validate:"required,email"`
	Password  string  `validate:"required,min=6"`
	Photos    []Photo `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
