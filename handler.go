package main

import (
	"context"
	"fmt"

	"github.com/josephus-git/gator/internal/database"
)

type command struct {
	Name    string
	Handler []string
}

type commands struct {
	Cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
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

func (c *commands) register(name string, f func(*state, command) error) {
	c.Cmds[name] = f
}
