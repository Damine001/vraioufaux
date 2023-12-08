package badge

import (
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BadgeRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {
	baseInstance := Database{DB: db}
	router.POST("/new", middleware.Authorize("badge", "write", enforcer), baseInstance.NewBadge)
	router.GET("/all", middleware.Authorize("badge", "read", enforcer), baseInstance.GetBadges)
}
