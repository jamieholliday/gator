package main

import (
	"context"
	"fmt"

	"github.com/jamieholliday/gator/internal/database"
)

func handlerGetFeedFollows(s *state, cmd command, user database.User) error {
	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("Error getting feeds: %w", err)
	}

	if len(feed_follows) != 0 {
		fmt.Println("Following: ")
	} else {
		fmt.Println("You are not following any feeds.")
	}

	for _, feed := range feed_follows {
		fmt.Printf("%s\n", feed.FeedsName.String)
	}

	return nil
}
