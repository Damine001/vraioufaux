package user

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
	"vertoufaux/api/app/role"
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new user
func (db Database) NewUser(ctx *gin.Context) {

	// init vars
	var user User
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(user.Name) || empty_reg.MatchString(user.Email) || empty_reg.MatchString(user.Password) || empty_reg.MatchString(user.Role) || empty_reg.MatchString(user.FamilyName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check role exists
	if check := role.GetRole(db.DB, user.Role); check != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this role does not exist"})
		return
	}

	// get values from session

	// hash password
	HashPassword(&user.Password)

	// init new user
	new_user := User{
		Name:       user.Name,
		FamilyName: user.FamilyName,
		Email:      user.Email,
		Password:   user.Password,
		LastLogin:  time.Now(),
	}

	// create user
	if _, err := NewUser(db.DB, new_user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// permission
	db.Enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(user.ID), 10), user.Role)

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all users from database
func (db Database) GetUsers(ctx *gin.Context) {

	// get users
	users, err := GetUsers(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// search users from database
func (db Database) SearchUsers(ctx *gin.Context) {

	// init vars
	var user User

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search users from database
	users, err := SearchUsers(db.DB, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (db Database) UpdateUser(ctx *gin.Context) {

	// init vars
	var user User
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	user_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(user.Name) || empty_reg.MatchString(user.Email) || empty_reg.MatchString(user.Password) || empty_reg.MatchString(user.Role) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check role exists
	if check := role.GetRole(db.DB, user.Role); check != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this role does not exist"})
		return
	}

	// hash password
	HashPassword(&user.Password)

	// ignore key attributs
	user.ID = uint(user_id)

	// update user
	if err = UpdateUser(db.DB, user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteUser(ctx *gin.Context) {

	// get id from path
	user_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete user
	if err = DeleteUser(db.DB, uint(user_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// get all users from database
func (db Database) GetUsersByRank(ctx *gin.Context) {

	// get users
	users, err := GetUsersByRank(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (db Database) me(ctx *gin.Context) {
	session := middleware.ExtractTokenValues(ctx)

	// Set the appropriate CORS headers
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.JSON(http.StatusOK, session.UserID)

}

func (db Database) GetUserByID(ctx *gin.Context) {

	// get id value from path
	user_id1, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user1, err := GetUserByID(db.DB, uint(user_id1))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user1)
}
