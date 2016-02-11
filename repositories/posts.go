package repositories

import (
	"models"

	"golang.org/x/net/context"
)

func IndexPosts() []models.SummarizedPost {
	return []models.SummarizedPost{
		models.SummarizedPost{Title: "Hello, World", Slug: "hello-world", Summary: "We want to try out Go!", AuthorId: "j123"},
		models.SummarizedPost{Title: "Hello, Go", Slug: "hello-go", Summary: "We want to try out the World!", AuthorId: "m321"},
		models.SummarizedPost{Title: "Mock DB", Slug: "mock-db", Summary: "This is just the mock database.", AuthorId: "jma"},
	}
}

func GetPost(slug string) models.ExpandedPost {
	tags := []models.Tag{models.Tag{"golang"}, models.Tag{"datastore"}}
	post := models.SummarizedPost{Title: "Hello, World", Slug: slug, Summary: "We want to try out Go!", AuthorId: "j123", Tags: tags}
	return models.ExpandedPost{Data: post, Body: "We want to try out Go, because we've heard it is so much fun!"}
}

func StorePost(context.Context) string {
	return "well done"
}
