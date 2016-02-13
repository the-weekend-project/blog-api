package api

import (
	"controllers"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router := httprouter.New()

	router.GET("/posts", controllers.IndexPosts)
	router.GET("/posts/:slug", controllers.GetPost)

	router.GET("/authors", controllers.IndexAuthors)
	router.GET("/users/:username", controllers.GetUser)
	router.PUT("/users/:username", controllers.StoreUser)

	http.Handle("/", router)
}
