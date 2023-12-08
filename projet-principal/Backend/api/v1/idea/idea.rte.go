package idea

import (
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IdeaRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {
	baseInstance := Database{DB: db}
	router.POST("/new", middleware.Authorize("idea", "write", enforcer), baseInstance.NewIdea)
	router.GET("/all", middleware.Authorize("idea", "read", enforcer), baseInstance.GetIdeas)
}
