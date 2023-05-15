package models

import (
	"github.com/abjelosevic88/go-fiber-postgres/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint    `json:"id" gorm:"primary key;autoIncrement"`
	Name     *string `json:"name" validate:"required,min=3,max=32"`
	Username *string `json:"username" validate:"required,min=3,max=32"`
	Password *string `json:"password" validate:"required,min=3,max=5"`
	Books    []Book  `json:"books"`
}

func (b *User) Validate() []*utils.ErrorResponse {
	return utils.ValidateModel(b)
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
