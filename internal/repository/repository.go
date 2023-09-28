package repository

import (
	"FreeMusic/internal/models"
	"golang.org/x/net/context"
)

type FileStorage interface {
	UploadFile(ctx context.Context, req models.UploadFileRequest) (*models.UploadFileResponse, error)
	DownloadFile(ctx context.Context, req models.DownloadFileRequest, fileExtension models.FileExtension) (*models.DownloadFileResponse, error)
	//DeleteFile() error
}

type Repository struct {
	FileStorage *FileStorage
}

func NewRepository(fileStorage FileStorage) *Repository {
	return &Repository{
		FileStorage: &fileStorage,
	}
}
