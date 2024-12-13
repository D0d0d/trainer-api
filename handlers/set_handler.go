package handlers

import (
	"net/http"
	"trainer-api/db"
	"trainer-api/models"

	"github.com/gin-gonic/gin"
)

// CreateSet создает новый набор вопросов
func CreateSet(c *gin.Context) {
	var set models.Set
	if err := c.ShouldBindJSON(&set); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&set).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать набор"})
		return
	}
	c.JSON(http.StatusCreated, set)
}

// GetSetByID возвращает набор вопросов по ID
func GetSetByID(c *gin.Context) {
	id := c.Param("id")
	var set models.Set
	if err := db.DB.Preload("Questions").First(&set, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Набор не найден"})
		return
	}
	c.JSON(http.StatusOK, set)
}

// AddQuestionToSet добавляет вопрос в набор
func AddQuestionToSet(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, существует ли набор
	var set models.Set
	if err := db.DB.First(&set, question.SetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Набор не найден"})
		return
	}

	if err := db.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить вопрос"})
		return
	}
	c.JSON(http.StatusCreated, question)
}
