package db

import "trainer-api/models"

// MigrateDB выполняет миграцию базы данных
func MigrateDB() {
	err := DB.AutoMigrate(&models.Set{}, &models.Question{})
	if err != nil {
		panic("Миграция не удалась: " + err.Error())
	}
}
