package models

type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Source      string    `json:"source"`
	State       string    `json:"state"`
	Category    string    `json:"category"`
	URL         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
	CollectedAt time.Time `json:"collected_at"`
	Content     string    `json:"content"`
}