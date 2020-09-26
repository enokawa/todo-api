package usecase

import (
	"github.com/enokawa/todo-api/domain"
	"time"
)

// ToDoUsecase is ...
type ToDoUsecase interface {
	Create(title string) (domain.ToDo, error)
}

type todoUsecase struct{}

var todoList = map[string]domain.ToDo{}

// NewToDoUsecase is ...
func NewToDoUsecase() ToDoUsecase {
	return todoUsecase{}
}

func (t todoUsecase) Create(title string) (domain.ToDo, error) {
	id := "xxxxxxxxxx"
	createdAt := time.Now()

	todo := domain.ToDo{ID: id, Title: title, CreatedAt: createdAt}

	todoList[id] = todo

	return todo, nil
}
