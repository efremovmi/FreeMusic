package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {

}

func newOkResponse(c *gin.Context, response http.Response) {
	c.JSON(http.StatusOK, response)
}
