package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: ./gator command <name>")
		os.Exit(1)
	}
	userName := cmd.Handler[1]

	// Check if user with that name already exists
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		fmt.Printf("Error: User '%s' does not exists.\n", userName)
		os.Exit(1)
	}

	err = s.cfg.SetUser(cmd.Handler[1])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Println("user has been set")
	return nil
}
