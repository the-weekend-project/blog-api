package models

import (
	"time"
)

type Post struct {
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	PostedAt  time.Time `json:"postedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    User      `json:"user"`
	Category  Category	`json:"category"`
	Tags	  []Tag		`json:"tags"`
}
