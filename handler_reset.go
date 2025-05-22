package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteAllUsers(context.Background())

	if err != nil {
		return fmt.Errorf("Error resetting db: %w", err)
	}

	fmt.Printf("Db reset")

	return nil
}
