package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jamieholliday/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := "2"
	if len(cmd.Args) > 0 {
		limit = cmd.Args[0]
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return fmt.Errorf("invalid limit: %s, must be a number", limit)
	}

	posts, err := s.db.GetFeedPostsForUser(context.Background(), database.GetFeedPostsForUserParams{
		Limit:  int32(limitInt),
		UserID: user.ID,
	})

	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s, URL: %s, Published At: %s\n", post.Title, post.Url, post.PublishedAt)
	}

	return nil
}
