package models

const (
	UserReader uint8 = 1
	UserAuthor uint8 = 3
	UserAdmin uint8 = 5
)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `datastore:",noindex"json:"email"`
	Password  string `datastore:",noindex" json:"-"`
	Level     uint8  `json:"-"`
	Website   string `datastore:",noindex" json:"website,omitempty"`
	Moto      string `datastore:",noindex" json:"moto,omitempty"`
}
