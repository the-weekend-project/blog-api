package controllers

import (
	"models"
	"repositories"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
)

// Gets all authors from the user's repository and prints them out as json
func IndexAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := appengine.NewContext(r)
	authors, err := repositories.GetUsersAboveLevel(models.UserAuthor, ctx)

	response, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Gets a user with the given username from the repository,
// or returns a 404 error if not found.
func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := appengine.NewContext(r)
	username := params.ByName("username")

	user, err := repositories.GetUser(username, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Creates a user with the given username, so long as that username is not already taken
func StoreUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	ctx := appengine.NewContext(r)

	user, err := repositories.GetUser(username, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = repositories.StoreUser(user, ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
