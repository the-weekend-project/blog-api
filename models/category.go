package models

type Category struct {
    Name  string `json:"name"`
    Description string `json:"description,omitempty"`
}
