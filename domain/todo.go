package domain

import (
	"time"
)

// ToDo struct
type ToDo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
