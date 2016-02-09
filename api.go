package api

import (
	"github.com/julienschmidt/httprouter"
	
	"net/http"

	"controllers"
)

func init() {
	router := httprouter.New()
    router.GET("/posts", controllers.IndexPosts)
    router.GET("/posts/:slug", controllers.GetPost)

    http.Handle("/", router)
}