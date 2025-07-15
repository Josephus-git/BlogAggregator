package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/josephus-git/gator/internal/database"
)

func registerUser(s *state, cmd command) error {
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: ./gator command <name>")
		os.Exit(1)
	}
	userName := cmd.Handler[1]
	currentTime := time.Now()

	// Check if user with that name already exists
	_, err := s.db.GetUser(context.Background(), userName)
	if err == nil {
		fmt.Printf("Error: User '%s' already exists.\n", userName)
		os.Exit(1)
	}

	// Create a new user
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      userName,
	}

	// Create the user in the database
	user, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("error in create user: %v", err)
	}
	err = s.cfg.SetUser(cmd.Handler[1])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Printf("user '%s' created successfully!\n", user.Name)
	log.Printf("User data: %v+\n", newUser) // log user's data for debugging
	return nil
}
