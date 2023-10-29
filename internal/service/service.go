package service

import (
	"FreeMusic/internal/models"
	"FreeMusic/internal/repository"
)

type FileManager interface {
	UploadFile(req models.UploadFileRequest) (*models.UploadFileResponse, error)
	DownloadFile(req models.DownloadFileRequest, fileExtension models.FileExtension) (*models.DownloadFileResponse, error)
	DownloadAudioImageFile(req models.DownloadFileRequest) (*models.DownloadAudioImageFileResponse, error)
	GetAllMusicFilesInfo(userID uint64) (*models.GetAllMusicFilesInfoResponse, error)
	DropFile(req models.DropFileRequest) error
}

type Service struct {
	FileManager
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		FileManager: NewFileManager(repos.FileStorage),
	}
}
