package main

import (
	"context"
	"fmt"
)

func users(s *state, cmd command) error {
	// get list of users in db
	names, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error in getting users: %v", err)
	}
	for _, name := range names {
		if name == s.cfg.Current_user_name {
			fmt.Printf("* %s (current)\n", name)
		} else {
			fmt.Printf("* %s\n", name)
		}
	}
	return nil
}
