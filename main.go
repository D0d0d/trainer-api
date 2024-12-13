package main

import (
	"github.com/joho/godotenv"
	"log"
	"trainer-api/db"
	"trainer-api/routes"
)

func main() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Println("Нет файла .env, используются стандартные переменные.")
	}

	// Подключение к БД
	db.ConnectDB()
	db.MigrateDB()

	// Настройка маршрутов
	r := routes.SetupRouter()

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}
}
