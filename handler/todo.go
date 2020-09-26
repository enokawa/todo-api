package handler

import (
	"fmt"
	"net/http"
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
	fmt.Fprint(w, "Create!\n")
}