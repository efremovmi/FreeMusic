package v1

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	appError "FreeMusic/internal/app_errors"
	"FreeMusic/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UploadFile @Summary UploadFile
// @Tags FileStorage
// @Description upload file
// @Accept multipart/form-data
// @Produce  json
//
// @Param Authorization header string true "Auth header"
// @Param filename formData string true "file name"
// @Param body formData file true "File to upload"
//
// @Success 200 {object} models.UploadFileResponse
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /file/upload [post]
func (h *Handler) uploadFile(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		logrus.Errorf("uploadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't get userID"})
		return
	}

	fileHeader, err := c.FormFile("body") // "file" - это имя поля формы для загрузки файла
	if err != nil {
		logrus.Errorf("uploadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't file"})
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)

	filename := c.PostForm("filename")
	if len(filename) == 0 {
		logrus.Errorf("uploadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't get filename"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		logrus.Errorf("uploadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't open file"})
		return
	}

	defer file.Close()

	req := models.UploadFileRequest{
		File:          file,
		FileName:      filename,
		FileExtension: fileExtension,
		UserID:        userID,
	}

	resp, err := h.services.UploadFile(req)
	if err != nil {
		logrus.Errorf("UploadFile err: %v", err)

		var appError *appError.DuplicateFound
		if errors.As(err, &appError) {
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse{"file with that name already exists"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't save file"})
		}
		return
	}

	c.JSON(http.StatusOK, resp)

}

// DownloadFile @Summary DownloadFile
// @Tags FileStorage
// @Description download file
// @Accept multipart/form-data
// @Produce  json
//
// @Param Authorization header string true "Auth header"
// @Param filename formData string true "file name"
//
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /file/download [post]
func (h *Handler) downloadFile(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		logrus.Errorf("downloadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't get userID"})
		return
	}

	filename := c.PostForm("filename")
	if len(filename) == 0 {
		logrus.Errorf("downloadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't get filename"})
		return
	}

	req := models.DownloadFileRequest{
		FileName: filename,
		UserID:   userID,
	}

	resp, err := h.services.DownloadFile(req, models.Any)
	if err != nil {
		logrus.Errorf("downloadFile err: %v", err)

		var appError *appError.FileNotFound
		if errors.As(err, &appError) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"file is not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't download file"})
		}

		return
	}
	defer resp.FileStream.Close()

	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", resp.FileName))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")

	_, err = io.Copy(c.Writer, resp.FileStream)
	if err != nil {
		logrus.Errorf("downloadFile err: can't copy file to resp stream, err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't download file"})
		return
	}

	c.Writer.WriteHeader(200)
}

// StreamAudio @Summary StreamAudio
// @Tags FileStorage
// @Description stream audio file
// @Accept multipart/form-data
// @Produce  json
//
// @Param Authorization header string true "Auth header"
// @Param filename formData string true "file name"
//
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /file/stream_audio [post]
func (h *Handler) streamAudio(c *gin.Context) {
	// TODO: доделать
	userID, err := getUserId(c)
	if err != nil {
		logrus.Errorf("downloadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't get userID"})
		return
	}

	filename := c.PostForm("filename")
	if len(filename) == 0 {
		logrus.Errorf("downloadFile err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't get filename"})
		return
	}

	req := models.DownloadFileRequest{
		FileName: filename,
		UserID:   userID,
	}

	resp, err := h.services.DownloadFile(req, models.MP3)
	if err != nil {
		logrus.Errorf("downloadFile err: %v", err)

		var appError *appError.FileNotFound
		if errors.As(err, &appError) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"file is not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"can't download file"})
		}

		return
	}
	defer resp.FileStream.Close()

	c.Header("Content-Type", "audio/mpeg")
	c.Status(http.StatusOK)

	bufferSize := 1024
	buffer := make([]byte, bufferSize)
	for {
		bytesRead, err := resp.FileStream.Read(buffer)
		if err != nil {
			break
		}
		c.Writer.Write(buffer[:bytesRead])
		c.Writer.Flush()
	}
}

// DropFile @Summary DropFile
// @Tags FileStorage
// @Description drop file
// @Accept  json
// @Produce  json
// @Param group_id   path int true "Group ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /file/drop [post]
func (h *Handler) dropFile(c *gin.Context) {

}
