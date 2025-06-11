package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/josephus-git/BlogAggregator/internal/database"
)

func browse(s *state, cmd command) error {
	limit := 2
	if len(cmd.Handler) > 1 {
		// Try to parse the string argument to an integer
		parsedLimit, err := strconv.Atoi(cmd.Handler[1])
		if err != nil {
			fmt.Println("Usage: ./BlogAggregator browse limit(optional integer)")
			os.Exit(1)
		}
		limit = parsedLimit
	}

	//get posts
	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("error getting posts: %v", err)
	}
	println("Posts:")

	for _, post := range posts {
		postStruct := reflect.ValueOf(post)
		for i := range postStruct.NumField() {
			field := postStruct.Type().Field(i)
			fmt.Printf("Title: %s, Value: %v\n", field.Name, postStruct.Field(i).Interface())
		}
		println("--------")
		println("")
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
	markParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:     time.Now(),
		ID:            fetchedFeed.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), markParams)
	if err != nil {
		return fmt.Errorf("error marking fetched feed: %v", err)
	}

	// fetch feed
	feed, err := s.db.Getfeed(context.Background(), fetchedFeed.Url)
	if err != nil {
		return fmt.Errorf("error in getting feed: %v", err)
	}

	// save post
	newPost := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   feed.CreatedAt,
		UpdatedAt:   feed.UpdatedAt,
		Title:       feed.Name,
		Url:         feed.Url,
		Description: feed.Name,
		PublishedAt: feed.LastFetchedAt.Time,
		FeedID:      feed.ID,
	}

	post, err := s.db.CreatePost(context.Background(), newPost)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			fmt.Println("posts updated")
		} else {
			log.Printf("error creating post: %v", err)
		}
	} else {
		fmt.Printf("successfully created post: %s\n", post.Title)
	}

	return nil
}
