package models

const (
	UserReader int8 = 1
	UserAuthor int8 = 3
	UserAdmin  int8 = 5
)

type User struct {
	Username  string `datastore:"-", json:"username"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email"`
	Password  []byte `datastore:",noindex" json:"-"`
	Level     int8   `json:"-"`
	Website   string `datastore:",noindex" json:"website,omitempty"`
	Moto      string `datastore:",noindex" json:"moto,omitempty"`
}
