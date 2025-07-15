package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

// browse retrieves and displays a specified number of the latest posts for the current user.
func browse(s *state, cmd command) error {
	limit := 2
	if len(cmd.Handler) > 1 {
		// Try to parse the string argument to an integer
		parsedLimit, err := strconv.Atoi(cmd.Handler[1])
		if err != nil {
			fmt.Println("Usage: ./gator browse limit(optional integer)")
			os.Exit(1)
		}
		limit = parsedLimit
	}

	//get posts
	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("error getting posts: %v", err)
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), s.cfg.Current_user_name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.Title)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
