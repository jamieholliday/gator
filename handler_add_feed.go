package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jamieholliday/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		// remturn an error
		return fmt.Errorf("usage: %s <name>, <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return fmt.Errorf("Could not get current user %s", s.cfg.CurrentUserName)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), UserID: user.ID, Name: name, Url: url, CreatedAt: time.Now(), UpdatedAt: time.Now()})

	if err != nil {
		return fmt.Errorf("Error creating feed: %w", err)
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

	fmt.Printf("Feed created: %s, %s, %s, %s, %s, %s\n", feed.ID, feed.UserID, feed.Name, feed.Url, feed.CreatedAt, feed.UpdatedAt)

	return nil
}
