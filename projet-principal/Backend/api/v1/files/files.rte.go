package files

import (
	"vertoufaux/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesAttachment(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/:id", middleware.Authorize("files", "write", enforcer), baseInstance.PostFiles) //upload

	router.GET("/download/:fileName/:id", middleware.Authorize("files", "read", enforcer), baseInstance.DownloadFile) //download

	router.GET("/getmostrecentfile", middleware.Authorize("files", "read", enforcer), baseInstance.GetMostRecentFile) //search most recent file
	
	router.GET("/all", middleware.Authorize("files", "read", enforcer), baseInstance.GetMostRecentFile)               //search most recent file

	/*
		router.DELETE("/:id", middleware.Authorize("files", "read", enforcer), baseInstance.DeleteFile)                   //delete file by id
		router.GET("/all/:id", middleware.Authorize("files", "read", enforcer), baseInstance.GetFilesInProject)           //get all files in a project id
	*/

}
