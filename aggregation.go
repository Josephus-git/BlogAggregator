package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
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

func scrapeFeeds(s *state) error {

	// fetch next feed id
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching nextfeed: %v", err)
	}

	//mark feed as fetched
	markParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:     time.Now(),
		ID:            feed.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), markParams)
	if err != nil {
		return fmt.Errorf("error marking fetched feed: %v", err)
	}

	// fetch feed
	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error in getting feed: %v", err)
	}
	for _, item := range fetchedFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
		})
	}

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		}
		log.Printf("Couldn't create post: %v", err)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(fetchedFeed.Channel.Item))

	return nil
}
