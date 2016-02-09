package models

import (
	"time"
)

type Post struct {
	Id        uint32    `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	PostedAt  time.Time `json:"postedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
