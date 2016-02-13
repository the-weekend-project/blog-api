package models

import (
	"time"
)

type Comment struct {
	Content   string    `datastore:",noindex" json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    User      `datastore:",noindex" json:"user"`
	Replies   []Comment `datastore:",noindex" json:"replies"`
}
