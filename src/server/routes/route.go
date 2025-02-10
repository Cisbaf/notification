package routes

import (
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.POST("/notification", NotificationRoute)
	router.GET("/connect", ConnectQrCodeRoute)
	// router.POST("/sendFile", SendFileWpp)
	return router
}
