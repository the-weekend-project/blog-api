package api

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"

	"repositories"
)

func init() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/posts/:slug", GetPost),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/", api.MakeHandler())
}

func GetPost(w rest.ResponseWriter, r *rest.Request) {
	slug := r.PathParam("slug")
	post := repositories.Get(slug)

	w.WriteJson(post)
}
