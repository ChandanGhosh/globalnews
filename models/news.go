package models

import (
	"github.com/dustin/go-humanize"
	"time"
)

// News ...
type News struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Article ...
type Article struct {
	Source      Source      `json:"source"`
	Author      string      `json:"author"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	URLToImage  string      `json:"urlToImage"`
	PublishedAt time.Time   `json:"publishedAt"`
	Content     interface{} `json:"content"`
}

// GetHumanFriendlyPublishedDate resolves the publish date to human understandable format
func (a *Article) GetHumanFriendlyPublishedDate() string{
	return humanize.Time(a.PublishedAt)
}

// Source ..,
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
