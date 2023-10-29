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

func (f *fileManager) DownloadFile(req models.DownloadFileRequest, fileExtension models.FileExtension) (*models.DownloadFileResponse, error) {
	resp, err := f.repo.DownloadFile(context.Background(), req, fileExtension)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadFile error")
	}

	return resp, nil
}

func (f *fileManager) DownloadAudioImageFile(req models.DownloadFileRequest) (*models.DownloadAudioImageFileResponse, error) {
	resp, err := f.repo.DownloadAudioImageFile(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadAudioImageFile error")
	}

	return resp, nil
}

func (f *fileManager) GetAllMusicFilesInfo(userID uint64) (*models.GetAllMusicFilesInfoResponse, error) {
	resp, err := f.repo.GetAllMusicFilesInfo(context.Background(), userID)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadFile error")
	}

	return resp, nil
}

func (f *fileManager) DropFile(req models.DropFileRequest) error {
	err := f.repo.DropFile(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "DropFile error")
	}

	return nil
}
