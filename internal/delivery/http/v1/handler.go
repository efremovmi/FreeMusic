package v1

import (
	"FreeMusic/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "FreeMusic/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	//r := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//auth := router.Group("/auth")
	//{
	//	auth.POST("/sign-up", h.signUp)
	//	auth.POST("/sign-in", h.signIn)
	//}

	fileAPI := router.Group("/file", h.userIdentity)
	{
		fileAPI.POST("/upload", h.uploadFile)
		fileAPI.POST("/download", h.downloadFile)
		fileAPI.DELETE("/drop", h.dropFile)
	}

	return router
}
