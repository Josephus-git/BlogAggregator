package main

import (
	"context"
	"fmt"

	"github.com/josephus-git/gator/internal/database"
)

// command represents a single executable command with its name and associated handler arguments.
type command struct {
	Name    string
	Handler []string
}

// commands holds a collection of named functions that can handle specific commands.
type commands struct {
	Cmds map[string]func(*state, command) error
}

// run executes the handler function associated with the given command name.
func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.Cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("available commands: login, reset, register, users, addfeed, feeds, follow, following, unfollow, agg, browse")
	}
	return cmdFunc(s, cmd)
}

// middlewareLoggedIn creates a middleware that ensures a user is authenticated before executing the handler.
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

// register adds a new command and its handler function to the commands map.
func (c *commands) register(name string, f func(*state, command) error) {
	c.Cmds[name] = f
}
