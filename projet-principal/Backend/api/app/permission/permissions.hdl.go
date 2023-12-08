package app_permission

import (
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new permission
func (db Database) NewPermission(ctx *gin.Context) {

	// init vars
	var err error
	var permission CasbinRule
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check if the sent content is compatible with CasbinRule
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(permission.V0) || empty_reg.MatchString(permission.V1) || empty_reg.MatchString(permission.V2) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check the action passed by the user
	if permission.V2 != "read" && permission.V2 != "write" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "permission is invalid"})
		return
	}

	// check if the role exists
	role, err := CheckRoleExists(db.DB, permission.V0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check if role name is empty
	if role.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "role name is invalid"})
		return
	}

	// check if the role has the policy if not give it the policy
	if hasPolicy := db.Enforcer.HasPolicy(permission.V0, permission.V1, permission.V2); !hasPolicy {
		db.Enforcer.AddPolicy(permission.V0, permission.V1, permission.V2)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all permissions
func (db Database) GetPermissions(ctx *gin.Context) {

	// get permissions
	permissions, err := GetPermissions(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, permissions)
}

// get permission by id
func (db Database) GetPermissionByID(ctx *gin.Context) {

	// get id value from path
	permission_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get permission by id
	permission, err := GetPermissionByID(db.DB, uint(permission_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if permission.V0 == "" || permission.V1 == "" || permission.V2 == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	ctx.JSON(http.StatusOK, permission)
}

// update permission
func (db Database) UpdatePermission(ctx *gin.Context) {

	// init vars
	var err error
	var permission CasbinRule
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check if the sent content is compatible with CasbinRule
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(permission.V0) || empty_reg.MatchString(permission.V1) || empty_reg.MatchString(permission.V2) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check if the role exists
	role, err := CheckRoleExists(db.DB, permission.V0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check the action passed by the user
	if permission.V2 != "read" && permission.V2 != "write" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "permission is invalid"})
		return
	}

	// check if role name is empty
	if role.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "role name is invalid"})
		return
	}

	// get id value from path
	permission_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get the row of the query
	db_permission, err := GetPermissionByID(db.DB, uint(permission_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check if it is a policy role
	if permission.V0 == "" || permission.V1 == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "permission not found"})
		return
	}

	// update the policy
	if permission.V1 == db_permission.V1 {
		db.Enforcer.UpdatePolicy([]string{db_permission.V0, db_permission.V1, db_permission.V2}, []string{permission.V0, permission.V1, permission.V2})
		ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid edit"})
	}
}

// remove permission
func (db Database) RemovePermission(ctx *gin.Context) {

	// get id value from path
	permission_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get the row of the query
	permission, err := GetPermissionByID(db.DB, uint(permission_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check if it is a policy role
	if permission.V0 == "" || permission.V1 == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "permission not found"})
		return
	}

	_, err = db.Enforcer.RemovePolicy(permission.V0, permission.V1, permission.V2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
