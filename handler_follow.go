package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/josephus-git/gator/internal/database"
)

// follow creates a new feed follow entry for the authenticated user based on a given feed URL.
func follow(s *state, cmd command, user database.User) error {
	// Check if input is accurate
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: ./gator follow <url>")
		os.Exit(1)
	}
	url := cmd.Handler[1]
	currentTime := time.Now()

	// get feed id
	feed, err := s.db.Getfeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error in getting feed id: %v", err)
	}

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
