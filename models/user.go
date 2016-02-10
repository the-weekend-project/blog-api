package models

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Level     uint8  `json:"-"`
	Website   string `json:"website,omitempty"`
	Moto      string `json:"moto,omitempty"`
}
