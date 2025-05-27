package main

import (
	"context"
	"fmt"

	"github.com/jamieholliday/gator/internal/database"
)

func handlerDeleteFeedFollows(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)

	if err != nil {
		return fmt.Errorf("Error getting feed by URL %s: %v\n", url, err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Error deleting feed follow for user %s and feed %s: %v", user.Name, feed.Name, err)
	}

	return nil
}
