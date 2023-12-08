package quiz

import (
	"gorm.io/gorm"
)

type Quiz struct {
	ID       uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Question string `gorm:"column:question;not null" json:"question"`
	Reponse  int    `gorm:"column:reponse;not null" json:"reponse"`
	gorm.Model
}

// create new user
func NewQuiz(db *gorm.DB, quiz Quiz) (Quiz, error) {
	return quiz, db.Create(&quiz).Error
}

func GetQuiz(db *gorm.DB) (quiz []Quiz, err error) {
	return quiz, db.Find(&quiz).Error
}
