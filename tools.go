package main

import (
	"context"
	"fmt"

	"github.com/hrncacz/go-gator/internal/config"
	"github.com/hrncacz/go-gator/internal/rss"
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
		fmt.Println(item.Title)
	}
	return nil
}
