package app_permission

import (
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesPermissions(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("permissions", "write", enforcer), baseInstance.NewPermission)
	router.GET("/all", middleware.Authorize("permissions", "read", enforcer), baseInstance.GetPermissions)
	router.GET("/:id", middleware.Authorize("permissions", "read", enforcer), baseInstance.GetPermissionByID)
	router.PUT("/:id", middleware.Authorize("permissions", "write", enforcer), baseInstance.UpdatePermission)
	router.DELETE("/:id", middleware.Authorize("permissions", "write", enforcer), baseInstance.RemovePermission)
}
