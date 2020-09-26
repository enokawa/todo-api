package handler

import (
	"fmt"
	"net/http"
)

// Handler interfaces
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

// NewHandler is ...
func NewHandler() Handler {
	return handler{}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create!\n")
}