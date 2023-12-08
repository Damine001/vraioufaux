package role

import (
	"gorm.io/gorm"
)

type Role struct {
	ID        uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Name      string `gorm:"column:name;not null;unique" json:"name"`
	CreatedBy uint   `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

func NewRole(db *gorm.DB, role Role) error {
	return db.Create(&role).Error
}

func GetRoles(db *gorm.DB) (role []Role, err error) {
	return role, db.Find(&role).Error
}

func GetRole(db *gorm.DB, name string) error {
	role := Role{}
	return db.Where("name=?", name).First(&role).Error
}

func SearchRoles(db *gorm.DB, role Role) (roles []Role, err error) {
	return roles, db.Where(&role).Find(&roles).Error
}

func UpdateRole(db *gorm.DB, role Role) error {
	return db.Where("id=?", role.ID).Updates(&role).Error
}

func DeleteRole(db *gorm.DB, role_id uint) error {
	return db.Where("id=?", role_id).Delete(&Role{}).Error
}
