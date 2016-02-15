package blogApi

import (
	"github.com/the-weekend-project/blogApi/controllers"

	"github.com/julienschmidt/httprouter"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/posts", controllers.IndexPosts)
	router.GET("/posts/:slug", controllers.GetPost)

	router.GET("/authors", controllers.IndexAuthors)
	router.GET("/users/:username", controllers.GetUser)
	router.PUT("/users/:username", controllers.StoreUser)

	return router
}
