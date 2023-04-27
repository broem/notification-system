package model

import (
	"time"
)

type Notification struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}