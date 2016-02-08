package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	post := Post{}

	json.NewEncoder(w).Encode(post)
}

type Post struct {
	Id        uint32    `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	PostedAt  time.Time `json:"postedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
