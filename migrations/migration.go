package migrations

import (
	"todolist/models"

	"gorm.io/gorm"
)

func MigrateDB(DB *gorm.DB) error {
	err := DB.AutoMigrate(&models.User{}, &models.Todo{})
	return err
}
