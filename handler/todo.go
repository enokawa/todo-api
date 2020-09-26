package handler

import (
	"fmt"
	"net/http"

	"github.com/enokawa/todo-api/usecase"
)

// ToDoHandler interfaces
type ToDoHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct{}

// NewHandler is ...
func NewHandler() ToDoHandler {
	return todoHandler{}
}

func (h todoHandler) Create(w http.ResponseWriter, r *http.Request) {
	u := usecase.NewToDoUsecase()

	res, err := u.Create("my title")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
