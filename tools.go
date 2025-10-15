package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/hrncacz/go-gator/internal/config"
	"github.com/hrncacz/go-gator/internal/database"
	"github.com/hrncacz/go-gator/internal/rss"
	"github.com/lib/pq"
)

func scrapeFeeds(s *config.State) error {
	ctx := context.Background()
	feed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}
	data, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err = s.DB.MarkAsFetched(ctx, feed.ID); err != nil {
		fmt.Println(err)
		return err
	}
	for _, item := range data.Channel.Item {
		publishDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Println(err)
			return err
		}
		description := sql.NullString{
			Valid:  true,
			String: item.Description,
		}

		if err = s.DB.CreatePost(ctx, database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishDate,
			FeedID:      feed.ID,
		}); err != nil {
			if err, ok := err.(*pq.Error); ok {
				if err.Code == "23505" {
					continue
				}
			}
			os.Exit(1)
		}
	}
	return nil
}
