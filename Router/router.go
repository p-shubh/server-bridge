package router

import (
	handlers "server-bridge/Handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the router and routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/upload", handlers.UploadHandler)
	r.GET("/download", handlers.DownloadHandler)
	r.GET("/list", handlers.ListDirectoryHandler)

	return r
}
