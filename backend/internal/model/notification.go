package model

import (
	"time"
)

type Notification struct {
	ID        int64     `json:"id,omitempty" db:"id"`
	Title     string    `json:"title,omitempty" db:"title"`
	Message   string    `json:"message,omitempty" db:"message"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at"`
}