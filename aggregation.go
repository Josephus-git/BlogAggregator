package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/josephus-git/BlogAggregator/internal/database"
)

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
		return fmt.Errorf("error creating post: %v", err)
	}

	fmt.Printf("successfully created post: %s", post.Title)

	/*
		// print the field name and its value
		feedStruct := reflect.ValueOf(feed)
		for i := range feedStruct.NumField() {
			field := feedStruct.Type().Field(i)
			fmt.Printf("Title: %s, Value: %v\n", field.Name, feedStruct.Field(i).Interface())
		}
	*/
	return nil
}
