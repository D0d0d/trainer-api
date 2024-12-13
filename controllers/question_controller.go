package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your_project/models"
)

// AddQuestion добавляет новый вопрос
func AddQuestion(c *gin.Context) {
	var input struct {
		Question string `json:"question" binding:"required"`
		Answers  string `json:"answers" binding:"required"`
		Correct  string `json:"correct" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка, что Correct — это валидный JSON-массив
	var correctAnswers []int
	if err := json.Unmarshal([]byte(input.Correct), &correctAnswers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Correct must be a valid JSON array of integers"})
		return
	}

	question := models.Question{
		Question: input.Question,
		Answers:  input.Answers,
		Correct:  input.Correct,
	}

	if err := models.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": question})
}

func GetQuestion(c *gin.Context) {
	id := c.Param("id")
	var question models.Question

	if err := models.DB.First(&question, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve question"})
		}
		return
	}

	// Преобразуем Correct в массив
	var correctAnswers []int
	if err := json.Unmarshal([]byte(question.Correct), &correctAnswers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse correct answers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       question.ID,
		"question": question.Question,
		"answers":  question.Answers,
		"correct":  correctAnswers,
	})
}
