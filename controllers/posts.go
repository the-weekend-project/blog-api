package controllers

import (
	"github.com/the-weekend-project/blogApi/repositories"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, err := json.Marshal(repositories.IndexPosts())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	slug := params.ByName("slug")
	post := repositories.GetPost(slug)

	response, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
