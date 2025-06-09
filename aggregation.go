package main

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/josephus-git/BlogAggregator/internal/database"
)

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

	// print the field name and its value
	feedStruct := reflect.ValueOf(feed)
	for i := range feedStruct.NumField() {
		field := feedStruct.Type().Field(i)
		fmt.Printf("Title: %s, Value: %v\n", field.Name, feedStruct.Field(i).Interface())
	}
	return nil
}
