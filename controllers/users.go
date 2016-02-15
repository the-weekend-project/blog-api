package controllers

import (
	"github.com/the-weekend-project/blogApi/models"
	"github.com/the-weekend-project/blogApi/repositories"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
)

// Gets all authors from the user's repository and prints them out as json
func IndexAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := appengine.NewContext(r)
	authors, err := repositories.GetUsersAboveLevel(ctx, models.UserAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	user := &models.User{Username: params.ByName("username")}

	user, err := repositories.GetUser(ctx, user)
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
	user := &models.User{Username: params.ByName("username")}
	ctx := appengine.NewContext(r)

	_, err := repositories.GetUser(ctx, user)
	if err == nil {
		http.Error(w, "Username already exists!", http.StatusConflict)
		return
	}

	user = &models.User{
		Username:  username,
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		Email:     r.FormValue("email"),
	}
	err = repositories.StoreUser(ctx, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
