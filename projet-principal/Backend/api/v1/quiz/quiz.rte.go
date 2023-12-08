package quiz

import (
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func QuizRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {
	baseInstance := Database{DB: db}
	router.POST("/new", middleware.Authorize("quiz", "write", enforcer), baseInstance.NewQuiz)
	router.GET("/all", middleware.Authorize("quiz", "write", enforcer), baseInstance.GetQuizs)
	router.POST("/score", middleware.Authorize("quiz", "write", enforcer), baseInstance.CalculPoints)

}
