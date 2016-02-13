package models

type Category struct {
	Name        string `json:"name"`
	Description string `datastore:",noindex" json:"description,omitempty"`
}
