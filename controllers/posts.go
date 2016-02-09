package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"repositories"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, err := json.Marshal(repositories.Index())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	slug := params.ByName("slug")
	post := repositories.Get(slug)

	response, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
