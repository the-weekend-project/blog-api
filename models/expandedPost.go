package models

type ExpandedPost struct {
	Data     SummarizedPost `json:"data"`
	Body     string         `datastore:",noindex" json:"body"`
	Comments []Comment      `json:"comments"`
}
