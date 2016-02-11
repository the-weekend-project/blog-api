package controllers

import (
	"repositories"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
)

func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	ctx := appengine.NewContext(r)

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
