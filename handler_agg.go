package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"github.com/jamieholliday/gator/internal/database"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	request.Header.Set("User-Agent", "gator")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching feed: recieved status code %d", response.StatusCode)
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	feed := &RSSFeed{}
	err = xml.Unmarshal(b, feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling feed: %w", err)
	}

	// unescape HTML entities in item descriptions
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
	}

	return feed, nil

}

func parsePublishedDate(dateStr string) time.Time {
	// Try RFC1123 format first (standard RSS date format)
	t, err := time.Parse(time.RFC1123, dateStr)
	if err == nil {
		return t
	}

	// Try RFC1123Z format (with numeric timezone)
	t, err = time.Parse(time.RFC1123Z, dateStr)
	if err == nil {
		return t
	}

	// Try RFC822 format (another common format)
	t, err = time.Parse(time.RFC822, dateStr)
	if err == nil {
		return t
	}

	// If all parsing attempts fail, return current time
	fmt.Printf("Failed to parse date: %s\n", dateStr)
	return time.Now()
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil || err == sql.ErrNoRows {
		return fmt.Errorf("Error getting next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            nextFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("Error marking feed fetched: %w", err)
	}

	feed, err := FetchFeed(context.Background(), nextFeed.Url)

	if err != nil {
		return fmt.Errorf("Error getting feed by URL %s: %w", nextFeed.Url, err)
	}

	for _, item := range feed.Channel.Item {
		// Process each item in the feed
		fmt.Printf("Feed item: %s\n", item)

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			FeedID:      nextFeed.ID,
			Title:       item.Title,
			Description: item.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Url:         item.Link,
			PublishedAt: parsePublishedDate(item.PubDate),
		})
		if err != nil {
			// PostgreSQL error code 23505 is for unique_violation
			if err.Error() == "pq: duplicate key value violates unique constraint" || err.Error() == "ERROR: duplicate key value violates unique constraint" {
				fmt.Printf("Post with URL %s already exists, skipping...\n", item.Link)
				continue
			} else {
				return fmt.Errorf("Error creating post: %w", err)
			}
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		// remturn an error
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)

	for range ticker.C {
		fmt.Println("Scraping feeds...")
		scrapeFeeds(s)
	}

	return nil
}
