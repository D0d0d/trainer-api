package models

// Set представляет собой коллекцию вопросов
type Set struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Questions   []Question `gorm:"foreignKey:SetID" json:"questions"`
}
