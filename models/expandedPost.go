package models

type ExpandedPost struct {
	Data     SummarizedPost `json:"data"`
	Body     string         `json:"body"`
	Comments []Comment      `json:"comments"`
}
