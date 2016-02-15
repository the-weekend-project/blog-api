package repositories_test

import (
	"github.com/the-weekend-project/blogApi/models"
	"github.com/the-weekend-project/blogApi/repositories"
	"testing"
	"time"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

// Should return users that are level 3 or above
func TestIndexAuthors(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	key := []*datastore.Key{
		datastore.NewKey(ctx, "User", "dranton", 0, nil),
		datastore.NewKey(ctx, "User", "testertim", 0, nil),
		datastore.NewKey(ctx, "User", "rookierook", 0, nil),
	}

	seededUsers := []*models.User{
		&models.User{FirstName: "Anton", Level: 5},
		&models.User{FirstName: "Tim", Level: 3},
		&models.User{FirstName: "Rookie", Level: 1},
	}

	if _, err := datastore.PutMulti(ctx, key, seededUsers); err != nil {
		t.Fatal(err)
	}

	time.Sleep(300 * time.Millisecond)

	var authors []models.User
	authors, err = repositories.GetUsersAboveLevel(ctx, 3)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(authors) != 2 {
		t.Errorf("Got %d authors, want %d", len(authors), 2)
	}

	if authors[0].Username != "testertim" {
		t.Errorf("Got username %d, want %d", authors[1].Username, "testertim")
	}
}

// Should return the user with the given username, or throw an Entity not found error
func TestGetUser(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	key := datastore.NewKey(ctx, "User", "testertim", 0, nil)
	if _, err := datastore.Put(ctx, key, &models.User{FirstName: "Tim"}); err != nil {
		t.Fatal(err)
	}

	user := &models.User{Username: "nonexistent"}
	if err = repositories.GetUser(ctx, user); err == nil || err.Error() != "datastore: no such entity" {
		t.Errorf("Error: %v; want datastore: no such entity", err)
	}

	user.Username = "testertim"
	err = repositories.GetUser(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	if firstName, want := user.FirstName, "Tim"; firstName != want {
		t.Errorf("User First Name %d, want %d", firstName, want)
	}

	if username, want := user.Username, "testertim"; username != want {
		t.Errorf("Username %d, want %d", username, want)
	}
}

// Should store the given user, after encrypting its password
func TestStoreUser(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	user := &models.User{Username: "testertim", FirstName: "Tim", Email: "tim@example.com", Password: []byte("I <3 golang")}
	err = repositories.StoreUser(ctx, user)

	if err != nil {
		t.Fatal(err)
	}

	key := datastore.NewKey(ctx, "User", "testertim", 0, nil)
	var storedUser models.User
	err = datastore.Get(ctx, key, &storedUser)
	if err != nil {
		t.Fatal(err)
	}

	if storedUser.FirstName != user.FirstName {
		t.Errorf("FirstName %d, want %d", storedUser.FirstName, user.FirstName)
	}

	if storedUser.Username != "" {
		t.Errorf("Username %d should be derived from the key, not stored", storedUser.Username)
	}

	if string(storedUser.Password) == "I <3 golang" {
		t.Errorf("Password %d should be hashed and not %d", string(storedUser.Password), "I <3 golang")
	}
}
