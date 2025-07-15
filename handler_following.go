package main

import (
	"context"
	"fmt"

	"github.com/josephus-git/gator/internal/database"
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
