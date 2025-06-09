package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/josephus-git/BlogAggregator/internal/config"
	"github.com/josephus-git/BlogAggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func Handlerrun() {
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
	ncmds := commands{}
	ncmds.Cmds = make(map[string]func(*state, command) error)

	// ensure accurate input arguments
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		fmt.Println("Usage: go run . command <name_if_required>")
		os.Exit(1)
	}
	// initialize the command struct
	cmd := command{}
	cmd.Name = arguments[0]
	cmd.Handler = arguments

	ncmds.Register("login", handlerLogin)
	ncmds.Register("register", registerUser)
	ncmds.Register("reset", resetData)
	ncmds.Register("users", users)
	ncmds.Register("addfeed", middlewareLoggedIn(addfeed))
	ncmds.Register("feeds", feeds)
	ncmds.Register("follow", middlewareLoggedIn(follow))
	ncmds.Register("following", middlewareLoggedIn(following))
	ncmds.Register("unfollow", middlewareLoggedIn(unfollow))
	ncmds.Register("agg", agg)

	err = ncmds.Run(&newState, cmd)
	if err != nil {
		fmt.Printf("error at run: %v\n", err)
		os.Exit(1)
	}
}
