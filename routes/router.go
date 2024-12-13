package routes

import (
	"trainer-api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter настройка маршрутов
func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// Эндпоинты для наборов
		api.POST("/sets", handlers.CreateSet)
		api.GET("/sets/:id", handlers.GetSetByID)

		// Эндпоинты для вопросов
		api.POST("/questions", handlers.AddQuestionToSet)
	}

	return r
}
