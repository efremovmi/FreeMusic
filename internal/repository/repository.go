package repository

import (
	"FreeMusic/internal/models"
	"golang.org/x/net/context"
)

type FileStorage interface {
	UploadFile(ctx context.Context, req models.UploadFileRequest) (*models.UploadFileResponse, error)
	DownloadFile(ctx context.Context, req models.DownloadFileRequest, fileExtension models.FileExtension) (*models.DownloadFileResponse, error)
	DownloadAudioImageFile(ctx context.Context, req models.DownloadFileRequest) (*models.DownloadAudioImageFileResponse, error)
	GetAllMusicFilesInfo(ctx context.Context, userID uint64) (*models.GetAllMusicFilesInfoResponse, error)
	DropFile(ctx context.Context, request models.DropFileRequest) error
}

type Repository struct {
	FileStorage *FileStorage
}

func NewRepository(fileStorage FileStorage) *Repository {
	return &Repository{
		FileStorage: &fileStorage,
	}
}
