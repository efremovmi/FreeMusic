package v1

import "github.com/gin-gonic/gin"

// SaveFile @Summary SaveFile
// @Tags FileStorage
// @Description save file
// @Accept  json
// @Produce  json
// @Param group_id   path int true "Group ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /file/save [post]
func (h *Handler) safeFile(c *gin.Context) {

}

// DeleteFile @Summary DeleteFile
// @Tags FileStorage
// @Description delete file
// @Accept  json
// @Produce  json
// @Param group_id   path int true "Group ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /file/delete [post]
func (h *Handler) deleteFile(c *gin.Context) {

}
