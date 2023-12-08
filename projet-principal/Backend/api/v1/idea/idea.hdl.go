package idea

import (
	"net/http"
	"os"
	"regexp"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func (db Database) NewIdea(ctx *gin.Context) {

	// init vars
	var idea Idea
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&idea); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(idea.Question) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// init new user
	new_user := Idea{
		Question: idea.Question,
		Vote:     idea.Vote,
	}

	// create user
	if _, err := NewIdea(db.DB, new_user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//user_update, _ := user.GetUserByID(db.DB, uint(num))
	//user_update.Points = user_update.Points + 20

	//_ = user.UpdateUser(db.DB, user_update)

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all users from database
func (db Database) GetIdeas(ctx *gin.Context) {

	// get users
	users, err := GetIdeas(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
