package repositories

import (
	"github.com/the-weekend-project/blogApi/models"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const (
	kind string = "User"
)

// Gets all users above and including the given level
// Will need to be refactored if it might ever return more than a 100 results
func GetUsersAboveLevel(ctx context.Context, level int8) ([]models.User, error) {
	q := datastore.NewQuery(kind).Filter("Level >=", level).Limit(100)
	var users []models.User
	keys, err := q.GetAll(ctx, &users)

	if err == nil {
		for i := range users {
			users[i].Username = keys[i].StringID()
		}
	}

	return users, err
}

// Gets the user with the given username
func GetUser(ctx context.Context, user *models.User) error {
	key := getUserKey(ctx, user.Username)
	err := datastore.Get(ctx, key, user)

	return err
}

// Persists the user to the datastore
func StoreUser(ctx context.Context, user *models.User) error {
	key := getUserKey(ctx, user.Username)
	var err error
	user.Password, err = bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	_, err = datastore.Put(ctx, key, user)
	return err
}

// Helper function to get the key for the User entities with the given username
func getUserKey(ctx context.Context, username string) *datastore.Key {
	return datastore.NewKey(
		ctx,
		kind,
		username,
		0,
		nil,
	)
}
