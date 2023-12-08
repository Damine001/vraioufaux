package badge

import (
	"net/http"
	"os"
	"regexp"
	"vertoufaux/api/app/user"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func (db Database) NewBadge(ctx *gin.Context) {

	// init vars
	var badge Badge
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&badge); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(badge.Title) || empty_reg.MatchString(badge.Description) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	if exists := user.CheckUserExists(db.DB, uint(badge.User_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	// session informations

	// init new role

	user_info, err := user.GetUserByID(db.DB, uint(badge.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	new_badge := Badge{
		ID:          badge.ID,
		User_id:     badge.User_id,
		Title:       badge.Title,
		User_name:   user_info.Name,
		Description: badge.Description,
	}

	// create new role
	if _, err := NewBadge(db.DB, new_badge); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

func (db Database) GetBadges(ctx *gin.Context) {

	roles, err := GetMostRecentBadge(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}
