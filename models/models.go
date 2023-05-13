package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := MigrateBooks(db)
	if err != nil {
		return err
	}

	return nil
}
