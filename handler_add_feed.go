package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/josephus-git/gator/internal/database"
)

// addFeed creates a new feed and a corresponding feed follow for the authenticated user.
func addFeed(s *state, cmd command, user database.User) error {
	// Ensure accurate querry input
	if len(cmd.Handler) < 3 {
		fmt.Println("Usage: ./gator addfeed <name> <url>")
		os.Exit(1)
	}
	name := cmd.Handler[1]
	url := cmd.Handler[2]
	currentTime := time.Now()

	// Create new feed in database
	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}
	fmt.Println("feed created successfully")
	fmt.Println(feed.Name)

	// Create new feed follow
	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	FeedFollows, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}
	fmt.Printf("feed follow created successfully; FeedName: %s, UserName: %s\n", FeedFollows.FeedName, FeedFollows.UserName)

	return nil
}
