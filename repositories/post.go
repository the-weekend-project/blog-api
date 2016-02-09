package repositories

import (
    "models"
)

func Get(slug string) models.Post {
    return models.Post{Id: 3, Title: "Hello, World", Slug: slug, Content: "We want to try out Go!"}
}
