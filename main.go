package main

import (
	"net/http"
	"log"

	"github.com/enokawa/todo-api/handler"
)

func main() {
	h := handler.NewHandler()
	http.HandleFunc("/", h.Create)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

