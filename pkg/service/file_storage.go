package service

import "FreeMusic/pkg/repository"

type fileManager struct {
	repo repository.FileStorage
}

// NewFileManager ...
func NewFileManager(repo repository.FileStorage) *fileManager {
	return &fileManager{
		repo: repo,
	}
}

func (f *fileManager) SaveFile() error {
	return nil
}

func (f *fileManager) DeleteFile() error {
	return nil
}
