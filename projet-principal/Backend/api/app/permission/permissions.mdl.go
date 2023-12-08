package app_permission

import (
	"vertoufaux/api/app/role"

	"gorm.io/gorm"
)

type CasbinRule struct {
	ID uint   `json:"id"`
	V0 string `json:"role"`
	V1 string `json:"object"`
	V2 string `json:"action"`
}

func GetPermissions(db *gorm.DB) (permissions []CasbinRule, err error) {
	return permissions, db.Table("casbin_rule").Find(&permissions, "ptype=?", "p").Error
}

func GetPermissionByID(db *gorm.DB, id uint) (permission CasbinRule, err error) {
	return permission, db.Table("casbin_rule").Where("id = ? AND ptype = 'p'", id).Find(&permission).Error
}

func CheckRolePermissionExists(db *gorm.DB, role string) (role_exists role.Role, err error) {
	return role_exists, db.Table("roles").Where("name=?", role).Error
}

func CheckRoleExists(db *gorm.DB, name string) (role role.Role, err error) {
	return role, db.Table("roles").Where("name = ?", name).Find(&role).Error
}
