package sources

import (
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"

	"yourproject/models"
)

func FetchRSS(feedURL, source, state, category string) ([]models.Article, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	for _, item := range feed.Items {
		articles = append(articles, models.Article{
			ID:          uuid.NewString(),
			Title:       item.Title,
			Source:      source,
			State:       state,
			Category:    category,
			URL:         item.Link,
			PublishedAt: *item.PublishedParsed,
			CollectedAt: time.Now().UTC(),
			Content:     item.Description,
		})
	}

	return articles, nil
}
