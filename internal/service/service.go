package service

import (
	"FreeMusic/internal/models"
	"FreeMusic/internal/repository"
)

type FileManager interface {
	UploadFile(req models.UploadFileRequest) (*models.UploadFileResponse, error)
	DownloadFile(req models.DownloadFileRequest) (*models.DownloadFileResponse, error)
}

type Service struct {
	FileManager
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		FileManager: NewFileManager(repos.FileStorage),
	}
}