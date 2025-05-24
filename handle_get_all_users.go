package main

import (
	"context"
	"fmt"
)

func handlerGetAllUsers(s *state, cmd command) error {

	users, err := s.db.GetAllUsers(context.Background())

	if err != nil {
		return fmt.Errorf("Error getting users: %w", err)
	}

	for _, user := range users {
		currentUser := ""
		if s.cfg.CurrentUserName == user.Name {
			currentUser = "(current)"
		}
		fmt.Printf("* %s %s\n", user.Name, currentUser)
	}

	return nil
}
