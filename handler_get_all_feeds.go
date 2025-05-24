package main

import (
	"context"
	"fmt"
)

func handlerGetAllFeeds(s *state, _cmd command) error {

	feeds, err := s.db.GetAllFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("Error getting feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%s %s %s\n", feed.Name, feed.Url, feed.UserName.String)
	}

	return nil
}
