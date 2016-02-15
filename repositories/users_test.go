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
        key := datastore.NewKey(ctx, "User", "testertim", 0, nil)
        if _, err := datastore.Put(ctx, key, &models.User{FirstName: "Tim", Level: 3}); err != nil {
                t.Fatal(err)
        }

        time.Sleep(500 * time.Millisecond)

        var authors []models.User
		authors, err = repositories.GetUsersAboveLevel(3, ctx)
        if err != nil {
                t.Errorf("Error: %v", err)
        }

        if len(authors) != 1 {
                t.Errorf("Got %d authors, want %d", len(authors), 1)
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

        if _, err = repositories.GetUser("nonexistent", ctx); err == nil || err.Error() != "datastore: no such entity" {
                t.Errorf("Error: %v; want datastore: no such entity", err)
        }

        var user models.User

        user, err = repositories.GetUser("testertim", ctx)
        if err != nil {
                t.Fatal(err)
        }

        if firstName, want := user.FirstName, "Tim"; firstName != want {
                t.Errorf("User First Name %d, want %d", firstName, want)
        }
}