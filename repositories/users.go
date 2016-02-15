package repositories

import (
	"github.com/the-weekend-project/blogApi/models"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const (
	kind string = "User" 
)

// Gets all users above and including the given level
// Will need to be refactored if it might ever return more than a 100 results
func GetUsersAboveLevel(level int8, ctx context.Context) ([]models.User, error) {
	q := datastore.NewQuery(kind).Filter("Level >=", level).Limit(100)
	var users []models.User
	_, err := q.GetAll(ctx, &users)

	return users, err
}

// Gets the user with the given username
func GetUser(username string, ctx context.Context) (models.User, error) {
	var user models.User
	key := getUserKey(username, ctx)
	err := datastore.Get(ctx, key, &user)

	return user, err
}

// Persists the user to the datastore
func StoreUser(user *models.User, ctx context.Context) error {
	key := getUserKey(user.Username, ctx)
	_, err := datastore.Put(ctx, key, user)
	return err
}

// Helper function to get the key for the User entities with the given username
func getUserKey(username string, ctx context.Context) *datastore.Key {
	return datastore.NewKey(
		ctx,
		kind,
		username,
		0,
		nil,
	)
}
