package database

import (
	"fmt"
	"os"
	"strconv"
	permission "vertoufaux/api/app/permission"
	"vertoufaux/api/app/role"
	"vertoufaux/api/app/user"
	"vertoufaux/api/v1/badge"
	"vertoufaux/api/v1/files"
	"vertoufaux/api/v1/quiz"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

// auto migrate datbles
func _auto_migrate_tables(db *gorm.DB) {

	// auto migrate casbin table
	if err := db.Table("casbin_rule").AutoMigrate(&permission.CasbinRule{}); err != nil {
		panic(fmt.Sprintf("Error while creating casbin table: %v", err))
	}

	// auto migrate user, role & company tables
	if err := db.AutoMigrate(
		&user.User{},
		&role.Role{},
		&badge.Badge{},
		&files.File{},
		&quiz.Quiz{},
	); err != nil {
		panic(err)
	}
}

// auto create root user
func _create_root_user(db *gorm.DB, enforcer *casbin.Enforcer) {

	// init vars:
	// root
	var user_id uint
	root_role := role.Role{}
	root_user := user.User{}
	//root_company := company.Company{}
	// default role
	user_role := role.Role{}

	// create root role
	// check root role exists
	if check := db.Where("name=?", os.Getenv("DEFAULT_ROOT")).Find(&root_role); check.RowsAffected == 0 && check.Error == nil {

		// create role user
		db_role := role.Role{Name: os.Getenv("DEFAULT_ROOT")}

		if err := db.Create(&db_role).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the root role: %v", err))
		}
	}

	// create root user
	// check root user exists
	if check := db.Where("email=?", os.Getenv("DEFAULT_EMAIL")).Find(&root_user); check.RowsAffected == 0 && check.Error == nil {

		// create user
		db_user := user.User{Name: os.Getenv("DEFAULT_NAME"), Email: os.Getenv("DEFAULT_EMAIL"), Password: os.Getenv("DEFAULT_PASSWORD"), Role: os.Getenv("DEFAULT_ROOT")}
		user.HashPassword(&db_user.Password)

		if err := db.Create(&db_user).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the root user: %v", err))
		}

		// used to save user id to create company
		user_id = db_user.ID
	} else {

		// used to save user id to create company
		user_id = root_user.ID
	}

	// add policy
	enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(user_id), 10), os.Getenv("DEFAULT_ROOT"))

	// create user
	if check := db.Where("name=?", os.Getenv("DEFAULT_USER")).Find(&user_role); check.RowsAffected == 0 && check.Error == nil {

		// create role user
		db_role := role.Role{Name: os.Getenv("DEFAULT_USER")}

		if err := db.Create(&db_role).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the user role: %v", err))
		}
	}

	// add policy
	enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(0), 10), os.Getenv("DEFAULT_USER"))

	// create company
	// // check company exists
	// if check := db.Where("name=?", os.Getenv("DEFAULT_COMPANY_NAME")).Find(&root_company); check.RowsAffected == 0 && check.Error == nil {

	// 	// create company
	// 	db_company := company.Company{Name: os.Getenv("DEFAULT_COMPANY_NAME"), Email: os.Getenv("DEFAULT_COMPANY_EMAIL"), Phone: os.Getenv("DEFAULT_COMPANY_PHONE"), Address: os.Getenv("DEFAULT_COMPANY_ADDRESS"), ManagedBy: user_id, CreatedBy: user_id}
	// 	err := db.Create(&db_company).Error
	// 	if err != nil {
	// 		panic(fmt.Sprintf("[WARNING] error while creating the root company: %v", err))
	// 	}

	// 	// edit user to add company id
	// 	if check := db.Where("email=?", os.Getenv("DEFAULT_EMAIL")).Find(&root_user); check.RowsAffected == 1 && check.Error == nil {
	// 		root_user.CompanyId = db_company.ID
	// 		if update := db.Where("id=?", root_user.ID).Updates(&root_user); update.Error != nil {
	// 			panic(fmt.Sprintf("[WARNING] error while updating the root user: %v", update.Error))
	// 		}
	// 	}
	// }
}

func AutoMigrateDatabase(db *gorm.DB, enforcer *casbin.Enforcer) {

	// create tables
	_auto_migrate_tables(db)

	// create root
	_create_root_user(db, enforcer)
}
