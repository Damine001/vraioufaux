package quiz

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"vertoufaux/api/app/user"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func (db Database) NewQuiz(ctx *gin.Context) {

	// init vars
	var quiz Quiz
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&quiz); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(quiz.Question) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// init new user
	new_user := Quiz{
		Question: quiz.Question,
		Reponse:  quiz.Reponse,
	}

	// create user
	if _, err := NewQuiz(db.DB, new_user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all users from database
func (db Database) GetQuizs(ctx *gin.Context) {

	// get users
	users, err := GetQuiz(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (db Database) CalculPoints(ctx *gin.Context) {

	userID := ctx.Param("id")
	num, _ := strconv.Atoi(userID)
	users_point, err := user.GetUserByID(db.DB, uint(num))

	if exists := user.CheckUserExists(db.DB, uint(num)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	All_quiz, _ := GetQuiz(db.DB)

	for _, quiz := range All_quiz {
		if quiz.Reponse == 0 {
			users_point.Points = users_point.Points + 1

		}
	}

	var _ = user.UpdateUser(db.DB, users_point)
	// Continue with the rest of your code
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}
