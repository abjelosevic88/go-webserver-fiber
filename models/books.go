package models

import (
	"github.com/abjelosevic88/go-fiber-postgres/utils"
	"gorm.io/gorm"
)

type Books struct {
	ID        uint    `json:"id" gorm:"primary key;autoIncrement"`
	Author    *string `json:"author" validate:"required,min=3,max=32"`
	Title     *string `json:"title" validate:"required,min=3,max=5"`
	Publisher *string `json:"publisher" validate:"required,min=3,max=32"`
}

func (b *Books) Validate() []*utils.ErrorResponse {
	return utils.ValidateModel(b)
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
