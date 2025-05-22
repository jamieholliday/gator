package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jamieholliday/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		// remturn an error
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	// check if user already exists
	user, err := s.db.GetUserByName(context.Background(), name)

	if err == nil && user.Name != "" {
		return fmt.Errorf("user %s already exists", name)
	}

	if err != nil && err.Error() != "sql: no rows in result set" {
		return fmt.Errorf("error checking user: %w", err)
	}

	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), Name: name, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	s.cfg.SetUser(newUser.Name)

	fmt.Printf("User registerd with name %s\n", newUser.Name)

	return nil
}
