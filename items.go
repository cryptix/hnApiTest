package main

type Item struct {
	Author   string `json:"author"`
	Children []struct {
		Item
		ParentID float64 `json:"parent_id"`
		StoryID  float64 `json:"story_id"`
	} `json:"children"`
	CreatedAt  string  `json:"created_at"`
	CreatedAtI float64 `json:"created_at_i"`
	ID         float64 `json:"id"`
	Points     float64 `json:"points"`
	Text       string  `json:"text"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
	URL        string  `json:"url"`
}
