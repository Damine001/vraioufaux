package role

import (
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesRoles(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db}

	router.POST("/new", middleware.Authorize("roles", "write", enforcer), baseInstance.NewRole)
	router.GET("/all", middleware.Authorize("roles", "read", enforcer), baseInstance.GetRoles)
	router.POST("/search", middleware.Authorize("roles", "read", enforcer), baseInstance.SearchRoles)
	router.PUT("/:id", middleware.Authorize("roles", "write", enforcer), baseInstance.UpdateRole)
	router.DELETE("/:id", middleware.Authorize("roles", "write", enforcer), baseInstance.DeleteRole)
}
