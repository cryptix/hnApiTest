package main

type User struct {
	About           interface{} `json:"about"`
	Avg             float64     `json:"avg"`
	CommentCount    float64     `json:"comment_count"`
	CreatedAt       string      `json:"created_at"`
	CreatedAtI      float64     `json:"created_at_i"`
	ID              float64     `json:"id"`
	Karma           float64     `json:"karma"`
	ObjectID        string      `json:"objectID"`
	SubmissionCount float64     `json:"submission_count"`
	Username        string      `json:"username"`
}
