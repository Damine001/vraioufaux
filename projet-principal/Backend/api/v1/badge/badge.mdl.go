package badge

import (
	"gorm.io/gorm"
)

type Badge struct {
	ID          uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	User_id     uint   `gorm:"column:user_id;autoIncrement;primaryKey" json:"user_id"`
	Title       string `gorm:"column:title;autoIncrement;primaryKey" json:"title"`
	User_name   string `gorm:"column:user_name;autoIncrement;primaryKey" json:"user_name"`
	Description string `gorm:"column:description;autoIncrement;primaryKey" json:"description"`
	gorm.Model
}

// create new badge
func NewBadge(db *gorm.DB, badge Badge) (Badge, error) {
	return badge, db.Create(&badge).Error
}

func GetMostRecentBadge(db *gorm.DB) (badge Badge, err error) {
	return badge, db.Order("created_at desc").First(&badge).Error
}
