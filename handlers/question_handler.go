package handlers

import (
	"net/http"
	"trainer-api/db"
	"trainer-api/models"

	"github.com/gin-gonic/gin"
)

// CreateQuestion обработчик для создания вопроса
func CreateQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать вопрос"})
		return
	}
	c.JSON(http.StatusCreated, question)
}

// GetQuestions обработчик для получения всех вопросов
func GetQuestions(c *gin.Context) {
	var questions []models.Question
	if err := db.DB.Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить вопросы"})
		return
	}
	c.JSON(http.StatusOK, questions)
}

// GetQuestionByID обработчик для получения вопроса по ID
func GetQuestionByID(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := db.DB.First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Вопрос не найден"})
		return
	}
	c.JSON(http.StatusOK, question)
}
