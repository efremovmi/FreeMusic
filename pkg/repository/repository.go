package repository

type FileStorage interface {
	SaveFile() error
	DeleteFile() error
}

type Repository struct {
	FileStorage
}

func NewRepository() *Repository {
	return &Repository{
		//FileStorage: NewAuthPostgres(db),
		//TodoList:    NewTodoListPostgres(db),
		//TodoItem:    NewTodoItemPostgres(db),
	}
}
