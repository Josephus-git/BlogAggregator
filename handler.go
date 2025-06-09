package main

import (
	"context"
	"fmt"
	"os"
	"time"

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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// check if user is available in database
		if s.cfg.Current_user_name == "" {
			return fmt.Errorf("authentication required: no user available")
		}
		user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			return fmt.Errorf("authentication failed: error fetching user")
		}
		return handler(s, cmd, user)
	}
}

func agg(s *state, cmd command) error {
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: go run . command <time_in_seconds(numbersonly)>")
		os.Exit(1)
	}
	duration := fmt.Sprintf("%ss", cmd.Handler[1])
	time_between_reqs, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("error creating time between requests: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		fmt.Println("feeds completely printed")
		fmt.Println("----------------")
		fmt.Println("")
	}

}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.Cmds[name] = f
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
