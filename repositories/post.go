package repositories

import (
	"models"
)

func Index() []models.Post {
	return []models.Post{
		models.Post{Id: 3, Title: "Hello, World", Slug: "hello-world", Content: "We want to try out Go!"},
		models.Post{Id: 2, Title: "Hello, Go", Slug: "hello-go", Content: "We want to try out the World!"},
		models.Post{Id: 1, Title: "Mock DB", Slug: "mock-db", Content: "This is just the mock database."},
	}
}

func Get(slug string) models.Post {
	return models.Post{Id: 3, Title: "Hello, World", Slug: slug, Content: "We want to try out Go!"}
}
