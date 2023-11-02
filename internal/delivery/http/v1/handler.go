package v1

import (
	"net/http"

	"FreeMusic/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "FreeMusic/docs"
)

// Handler ...
type Handler struct {
	services *service.Service
}

// NewHandler ...
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes ...
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//auth := router.Group("/auth")
	//{
	//	auth.POST("/sign-up", h.signUp)
	//	auth.POST("/sign-in", h.signIn)
	//}

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	fileAPI := router.Group("/v1/file", h.userIdentity)
	{
		fileAPI.POST("/upload", h.uploadFile)
		fileAPI.POST("/download-audio", h.downloadAudio)
		fileAPI.POST("/download", h.downloadFile)
		fileAPI.POST("/download-audio-image", h.downloadAudioImage)
		fileAPI.GET("/get-all-music", h.getAllMusicFilesInfo)
		fileAPI.DELETE("/drop", h.dropFile)
	}

	return router
}
