package main

import (
	"context"
	"fmt"
)

// resetData clears all data from the database.
func resetData(s *state, cmd command) error {
	err := s.db.ResetTable(context.Background())
	if err != nil {
		return fmt.Errorf("error in reset Table: %v", err)
	}
	fmt.Println("The database has been reset")
	return nil
}
