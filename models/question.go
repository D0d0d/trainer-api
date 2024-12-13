package models

import "github.com/lib/pq"

// Question модель для вопросов
// Question представляет вопрос, связанный с набором
type Question struct {
	ID       uint           `gorm:"primaryKey" json:"id"`
	SetID    uint           `gorm:"not null" json:"set_id"` // Внешний ключ на набор
	Question string         `gorm:"type:text;not null" json:"question"`
	Answers  pq.StringArray `gorm:"type:text[];not null" json:"answers"` // JSON с массивом ответов
	Correct  pq.Int32Array  `gorm:"type:int[];not null" json:"correct"`  // Индекс правильного ответа
}
