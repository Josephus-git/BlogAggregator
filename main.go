package main

import (
	_ "github.com/lib/pq"

	"github.com/josephus-git/gator/internal/config"
	"github.com/josephus-git/gator/internal/database"
)

// state holds the application's current database queries and configuration.
type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	run()
}
