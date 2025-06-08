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

func following(s *state, cmd command, user database.User) error {

	feedsFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error in getting feed followed by user: %v", err)
	}

	fmt.Printf("user: %s follows:\n", s.cfg.Current_user_name)
	for _, feedFollow := range feedsFollowing {
		fmt.Println(feedFollow.FeedName)
	}
	return nil
}

func feeds(s *state, cmd command) error {
	allFeeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error in getting feeds: %v", err)
	}
	for _, feed := range allFeeds {
		fmt.Println(feed)
	}
	return nil
}

func follow(s *state, cmd command, user database.User) error {
	// Check if input is accurate
	if len(cmd.Handler) < 2 {
		fmt.Println("Usage: go run . follow <url>")
		os.Exit(1)
	}
	url := cmd.Handler[1]
	currentTime := time.Now()

	// get feed id
	feedID, err := s.db.GetfeedId(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error in getting feed id: %v", err)
	}

	// Create new feed follow
	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    feedID,
	}

	FeedFollows, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Printf("feed follow created successfully; FeedName: %s, UserName: %s\n", FeedFollows.FeedName, FeedFollows.UserName)

	return nil

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
	err := s.db.ResetTable(context.Background())
	if err != nil {
		return fmt.Errorf("error in reset Table: %v", err)
	}
	fmt.Println("The database has been reset")
	return nil
}

func addfeed(s *state, cmd command, user database.User) error {
	// Ensure accurate querry input
	if len(cmd.Handler) < 3 {
		fmt.Println("Usage: go run . addfeed <name> <url>")
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
