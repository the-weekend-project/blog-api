package repositories

import (
	"models"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func GetUser(username string, ctx context.Context) (models.User, error) {
	var user models.User
	key := getUserKey(user.Username, ctx)
	err := datastore.Get(ctx, key, &user)

	return user, err
}

func StoreUser(user models.User, ctx context.Context) error {
	key := getUserKey(user.Username, ctx)
	_, err := datastore.Put(ctx, key, user)
	return err
}

func getUserKey(username string, ctx context.Context) *datastore.Key {
	return datastore.NewKey(
		ctx,
		"User",
		username,
		0,
		nil,
	)
}
