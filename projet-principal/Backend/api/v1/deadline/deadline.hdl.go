package deadline

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

func (db Database) NewDeadline(ctx *gin.Context) {

	// init vars
	var deadline Deadline
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&deadline); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(deadline.Discription) || empty_reg.MatchString(deadline.ImageLien) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// session informations

	// init new role

	newDeadline := Deadline{
		ImageLien:   deadline.ImageLien,
		Titre:       deadline.Titre,
		Discription: deadline.Discription,
		Deadline:    deadline.Deadline,
	}

	// create new role
	if _, err := NewDeadline(db.DB, newDeadline); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

func (db Database) GetDeadlines(ctx *gin.Context) {

	roles, err := GetDeadline(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}
