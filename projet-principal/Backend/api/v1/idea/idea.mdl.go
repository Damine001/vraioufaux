package idea

import (
	"gorm.io/gorm"
)

type Idea struct {
	ID       uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Question string `gorm:"column:question;not null" json:"question"`
	Vote     int    `gorm:"column:vote;not null" json:"vote"`
	gorm.Model
}

// create new user
func NewIdea(db *gorm.DB, idea Idea) (Idea, error) {
	return idea, db.Create(&idea).Error
}

func GetIdeas(db *gorm.DB) (idea []Idea, err error) {
	return idea, db.Find(&idea).Error
}
