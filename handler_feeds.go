package main

import (
	"context"
	"fmt"
)

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
