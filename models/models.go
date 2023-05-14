package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := MigrateBooks(db); err != nil {
		return err
	}

	if err := MigrateUsers(db); err != nil {
		return err
	}

	return nil
}
