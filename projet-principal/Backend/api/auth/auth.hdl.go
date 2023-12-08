package app_auth

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
	"vertoufaux/api/app/user"
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// signup user
func (db Database) SignUpUser(ctx *gin.Context) {

	// init vars
	var account InsertUser
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// upmarshal sent json
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check field validity
	if empty_reg.MatchString(account.Name) || empty_reg.MatchString(account.Email) || empty_reg.MatchString(account.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// hash password
	user.HashPassword(&account.Password)

	// create new user
	new_user := user.User{
		Name:       account.Name,
		FamilyName: account.FamilyName,
		Email:      account.Email,
		Password:   account.Password,
		Points:     0,
		Role:       "user",
		LastLogin:  time.Now(), // Provide the appropriate value for last login
	}

	// create user
	saved_user, err := user.NewUser(db.DB, new_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// super add will add role
	db.Enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(saved_user.ID), 10), saved_user.Role)

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// signin user
func (db Database) SignInUser(ctx *gin.Context) {

	// init cars
	var user_login UserLogIn
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user_login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// check field validity
	if empty_reg.MatchString(user_login.Email) || empty_reg.MatchString(user_login.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check if email exists
	dbUser, err := user.GetUserByEmail(db.DB, user_login.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "no Such User Found"})
		return
	}

	// update last login
	dbUser.LastLogin = time.Now()

	// update user
	if err := user.UpdateUser(db.DB, dbUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// compare password
	if isTrue := user.ComparePassword(dbUser.Password, user_login.Password); isTrue {

		// generate token
		token := middleware.GenerateToken(dbUser.ID, dbUser.Role)
		ctx.JSON(http.StatusOK, UserLogedIn{Token: token})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "password not matched"})
}
