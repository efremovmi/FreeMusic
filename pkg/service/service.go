package service

import "FreeMusic/pkg/repository"

//type Service struct {
//	FileManager
//}
//
//func NewService(repos *repository.Repository) *Service {
//	return &Service{
//		Authorization: NewAuthService(repos.Authorization),
//		TodoList:      NewTodoListService(repos.TodoList),
//		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
//	}
//}

type FileManager interface {
	SaveFile() error
	DeleteFile() error
}

type Service struct {
	FileManager
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		FileManager: NewFileManager(repos.FileStorage),
	}
}
