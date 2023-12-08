package v1

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"vertoufaux/api/v1/badge"
	"vertoufaux/api/v1/deadline"
	"vertoufaux/api/v1/files"
	"vertoufaux/api/v1/idea"
	"vertoufaux/api/v1/quiz"
)

func RoutesV1(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	badge.BadgeRoutes(router.Group("/BadgeRoutes"), db, enforcer)
	deadline.DeadlineRoutes(router.Group("/deadline"), db, enforcer)
	files.RoutesAttachment(router.Group("/files"), db, enforcer)
	quiz.QuizRoutes(router.Group("/quiz"), db, enforcer)
	idea.IdeaRoutes(router.Group("/idea"), db, enforcer)

}
