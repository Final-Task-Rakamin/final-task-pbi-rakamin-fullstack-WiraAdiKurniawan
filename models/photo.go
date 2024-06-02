package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	Title    string `validate:"required"`
	Caption  string `validate:"required"`
	PhotoURL string `validate:"required"`
	UserID   uint
}
