package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/josephus-git/BlogAggregator/internal/database"
)

func browse(s *state, cmd command) error {
	limit := 2
	if len(cmd.Handler) > 1 {
		// Try to parse the string argument to an integer
		parsedLimit, err := strconv.Atoi(cmd.Handler[1])
		if err != nil {
			fmt.Println("Usage: go run . browse limit(optional integer)")
			os.Exit(1)
		}
		limit = parsedLimit
	}

	//get posts
	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("error getting posts: %v", err)
	}
	println("get here")

	for _, post := range posts {
		fmt.Println(post)
		postStruct := reflect.ValueOf(post)
		for i := range postStruct.NumField() {
			field := postStruct.Type().Field(i)
			fmt.Printf("Title: %s, Value: %v", field.Name, postStruct.Field(i).Interface())
		}
	}
	return nil
}

func scrapeFeeds(s *state) error {
	// fetch next feed id
	fetchedFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching nextfeed: %v", err)
	}

	//mark feed as fetched
	currentTime := time.Now()
	markParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: currentTime, Valid: true},
		UpdatedAt:     currentTime,
		ID:            fetchedFeed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markParams)
	if err != nil {
		return fmt.Errorf("error marking fetched feed: %v", err)
	}

	feed, err := s.db.Getfeed(context.Background(), fetchedFeed.Url)
	if err != nil {
		return fmt.Errorf("error in getting feed: %v", err)
	}

	// store posts
	postParams := CreatePostParams {
	ID:          uuid.New(),
	CreatedAt:   time.Time
	UpdatedAt:   time.Time
	Title:       string
	Url:         string
	Description: string
	PublishedAt: time.Time
	FeedID:      uuid.UUID
}

	// print the field name and its value
	feedStruct := reflect.ValueOf(feed)
	for i := range feedStruct.NumField() {
		field := feedStruct.Type().Field(i)
		fmt.Printf("Title: %s, Value: %v\n", field.Name, feedStruct.Field(i).Interface())
	}
	return nil
}
