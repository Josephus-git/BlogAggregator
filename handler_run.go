package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/josephus-git/gator/internal/config"
	"github.com/josephus-git/gator/internal/database"
)

// run initializes the application, sets up the database and configuration, and executes the requested command.
func run() {
	// obtain the db queries
	dURL := "postgres://postgres:postgres@localhost:5432/gator"
	db, err := sql.Open("postgres", dURL)
	if err != nil {
		fmt.Printf("dbError: %v\n", err)
		return
	}
	dbQueries := database.New(db)

	// obtain new config
	newConfig, err := config.Read()
	if err != nil {
		fmt.Printf("config Error: %v\n", err)
		return
	}

	// asign them to a new state
	newState := state{db: dbQueries, cfg: newConfig}
	newCommands := commands{}
	newCommands.Cmds = make(map[string]func(*state, command) error)

	// ensure accurate input arguments
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		fmt.Println("Usage: gator command <name_if_required>")
		os.Exit(1)
	}
	// initialize the command struct
	cmd := command{}
	cmd.Name = arguments[0]
	cmd.Handler = arguments

	newCommands.register("login", handlerLogin)
	newCommands.register("register", registerUser)
	newCommands.register("reset", resetData)
	newCommands.register("users", users)
	newCommands.register("addfeed", middlewareLoggedIn(addFeed))
	newCommands.register("feeds", feeds)
	newCommands.register("follow", middlewareLoggedIn(follow))
	newCommands.register("following", middlewareLoggedIn(following))
	newCommands.register("unfollow", middlewareLoggedIn(unfollow))
	newCommands.register("agg", aggregate)
	newCommands.register("browse", browse)

	err = newCommands.run(&newState, cmd)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
