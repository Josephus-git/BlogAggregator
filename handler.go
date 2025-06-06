package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/josephus-git/BlogAggregator/internal/database"
)

type command struct {
	Name    string
	Handler []string
}

type commands struct {
	Cmds map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	cmdFunc, ok := c.Cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	return cmdFunc(s, cmd)
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.Cmds[name] = f
}

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

func resetData(s *state, cmd command) error {
	//reset table in users db
	err := s.db.ResetTable(context.Background())
	if err != nil {
		return fmt.Errorf("error in create user: %v", err)
	}
	fmt.Println("The database has been reset")
	return nil
}

func registerUser(s *state, cmd command) error {
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: go run . command <name>")
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: go run . command <name>")
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
