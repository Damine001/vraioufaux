package role

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"vertoufaux/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// create new role
func (db Database) NewRole(ctx *gin.Context) {

	// init vars
	var role Role
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(role.Name) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// session informations
	session := middleware.ExtractTokenValues(ctx)

	// init new role
	new_role := Role{
		Name:      role.Name,
		CreatedBy: session.UserID,
	}

	// create new role
	if err := NewRole(db.DB, new_role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all roles
func (db Database) GetRoles(ctx *gin.Context) {

	// get roles from database
	roles, err := GetRoles(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

// search roles from database
func (db Database) SearchRoles(ctx *gin.Context) {

	// init vars
	var role Role

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search roles
	roles, err := SearchRoles(db.DB, role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

// update role
func (db Database) UpdateRole(ctx *gin.Context) {

	// init vars
	var role Role
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(role.Name) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get id value from path
	role_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// init update role
	update_role := Role{
		ID:   uint(role_id),
		Name: role.Name,
	}

	if err = UpdateRole(db.DB, update_role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete role from database
func (db Database) DeleteRole(ctx *gin.Context) {

	// get id value from path
	role_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete role
	err = DeleteRole(db.DB, uint(role_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
