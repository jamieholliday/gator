package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jamieholliday/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		// remturn an error
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting feed by URL %s: %v\n", url, err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("Error creating feed follow: %w", err)
	}

	fmt.Printf("Created feed follow for user %s and feed %s\n", user.Name, feed.Name)

	return nil
}
