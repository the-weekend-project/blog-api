package models

import (
	"time"
)

type SummarizedPost struct {
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Summary    string    `datastore:",noindex" json:"summary"`
	CreatedAt  time.Time `json:"createdAt"`
	PostedAt   time.Time `json:"postedAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	AuthorId   string    `json:"authorId"`
	CategoryId string    `json:"categoryId"`
	Tags       []Tag     `json:"tags"`
}
