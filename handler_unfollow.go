package main

import (
	"context"
	"fmt"
	"os"

	"github.com/josephus-git/gator/internal/database"
)

func unfollow(s *state, cmd command, user database.User) error {
	// Check if input is accurate
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: ./gator Unfollow <url>")
		os.Exit(1)
	}
	// get feed id
	url := cmd.Handler[1]
	feed, err := s.db.Getfeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed id: %s", err)
	}

	deleteparams := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), deleteparams)
	if err != nil {
		return fmt.Errorf("error deleting feed follow: %s", err)
	}

	fmt.Printf("Successfully deleted Feed: %s for User: %s", feed.Name, user.Name)

	return nil
}
