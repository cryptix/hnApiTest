package main

type Search struct {
	Hits []struct {
		HighlightResult struct {
			Author struct {
				MatchLevel   string        `json:"matchLevel"`
				MatchedWords []interface{} `json:"matchedWords"`
				Value        string        `json:"value"`
			} `json:"author"`
			StoryText struct {
				MatchLevel   string        `json:"matchLevel"`
				MatchedWords []interface{} `json:"matchedWords"`
				Value        string        `json:"value"`
			} `json:"story_text"`
			Title struct {
				MatchLevel   string   `json:"matchLevel"`
				MatchedWords []string `json:"matchedWords"`
				Value        string   `json:"value"`
			} `json:"title"`
			URL struct {
				MatchLevel   string        `json:"matchLevel"`
				MatchedWords []interface{} `json:"matchedWords"`
				Value        string        `json:"value"`
			} `json:"url"`
		} `json:"_highlightResult"`
		Tags        []string    `json:"_tags"`
		Author      string      `json:"author"`
		CommentText interface{} `json:"comment_text"`
		CreatedAt   string      `json:"created_at"`
		CreatedAtI  float64     `json:"created_at_i"`
		NumComments float64     `json:"num_comments"`
		ObjectID    string      `json:"objectID"`
		ParentID    interface{} `json:"parent_id"`
		Points      float64     `json:"points"`
		StoryID     interface{} `json:"story_id"`
		StoryText   string      `json:"story_text"`
		StoryTitle  interface{} `json:"story_title"`
		StoryURL    interface{} `json:"story_url"`
		Title       string      `json:"title"`
		URL         string      `json:"url"`
	} `json:"hits"`
	HitsPerPage      float64 `json:"hitsPerPage"`
	NbHits           float64 `json:"nbHits"`
	NbPages          float64 `json:"nbPages"`
	Page             float64 `json:"page"`
	Params           string  `json:"params"`
	ProcessingTimeMS float64 `json:"processingTimeMS"`
	Query            string  `json:"query"`
}
