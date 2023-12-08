package deadline

import (
	"time"

	"gorm.io/gorm"
)

type Deadline struct {
	ID          uint      `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	ImageLien   string    `gorm:"column:image_lien;not null" json:"image_lien"`
	Titre       string    `gorm:"column:titre;not null" json:"titre"`
	Discription string    `gorm:"column:discription;not null;unique" json:"discription"`
	Deadline    time.Time `gorm:"column:password;not null" json:"password"`
	gorm.Model
}

// create new user
func NewDeadline(db *gorm.DB, deadline Deadline) (Deadline, error) {
	return deadline, db.Create(&deadline).Error
}

func GetDeadline(db *gorm.DB) (deadline []Deadline, err error) {
	return deadline, db.Find(&deadline).Error
}
