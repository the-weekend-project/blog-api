package models

import (
	"time"
)

type Comment struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    User      `json:"user"`
	Category  Category  `json:"category"`
	Replies   []Comment `json:"replies"`
}
