package service

import (
	"FreeMusic/internal/models"
	"FreeMusic/internal/repository"
	"context"
	"github.com/pkg/errors"
)

type fileManager struct {
	repo repository.FileStorage
}

// NewFileManager ...
func NewFileManager(repo *repository.FileStorage) *fileManager {
	return &fileManager{
		repo: *repo,
	}
}

func (f *fileManager) UploadFile(req models.UploadFileRequest) (*models.UploadFileResponse, error) {
	resp, err := f.repo.UploadFile(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "UploadFile error")
	}

	return resp, nil
}

func (f *fileManager) DownloadFile(req models.DownloadFileRequest) (*models.DownloadFileResponse, error) {
	resp, err := f.repo.DownloadFile(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadFile error")
	}

	return resp, nil
}
