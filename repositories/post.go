package repositories

import (
	"models"
)

func Index() []models.Post {
	user := models.User{Username: "Athlos", FirstName: "Johnny", LastName: "Magrippis", Email: "j@magrippis.com", Website: "magrippis.com"}
	return []models.Post{
		models.Post{Title: "Hello, World", Slug: "hello-world", Content: "We want to try out Go!", Author: user},
		models.Post{Title: "Hello, Go", Slug: "hello-go", Content: "We want to try out the World!", Author: user},
		models.Post{Title: "Mock DB", Slug: "mock-db", Content: "This is just the mock database.", Author: user},
	}
}

func Get(slug string) models.Post {
	tags := []models.Tag{models.Tag{"golang"}, models.Tag{"datastore"}}
	user := models.User{Username: "Athlos", FirstName: "Johnny", LastName: "Magrippis", Email: "j@magrippis.com", Website: "magrippis.com"}
	return models.Post{Title: "Hello, World", Slug: slug, Content: "We want to try out Go!", Author: user, Category: models.Category{Name: "backend"}, Tags: tags}
}
