package main

type Article struct {
	ID string `json:"item_id"`

	Domain string `json:"domain"`

	GivenTitle string `json:"given_title"`
	GivenURL   string `json:"given_url"`

	ResolvedTitle string `json:"resolved_title"`
	ResolvedURL   string `json:"resolved_url"`

	WordCount string `json:"word_count"`
	Excerpt   string `json:"excerpt"`

	IsFavorite  string `json:"favorite"`
	IsArticle   string `json:"is_article"`
	HasVideo    string `json:"has_video"`
	HasImage    string `json:"has_image"`
	TopImageURL string `json:"top_image_url"`

	CreatedAt Timestamp `json:"time_added"`
	UpdatedAt Timestamp `json:"time_updated"`
	ReadAt    Timestamp `json:"time_read"`
}
