package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/hrncacz/go-gator/internal/command"
	"github.com/hrncacz/go-gator/internal/config"
	"github.com/hrncacz/go-gator/internal/database"
)

func middlewareLoggedIn(handler func(state *config.State, cmd command.Command, user database.User) error) func(*config.State, command.Command) error {
	return func(s *config.State, cmd command.Command) error {
		user, err := s.DB.SelectUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return errors.New("user not found")
		}
		return handler(s, cmd, user)
	}
}

func handlerLogin(s *config.State, cmd command.Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("missing arguments for login command")
	} else if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	}

	name := cmd.Args[0]
	users, err := s.DB.SelectUser(context.Background(), name)
	if err != nil {
		os.Exit(1)
	}
	s.Cfg.SetUser(users.Name)
	return nil
}

func handlerRegister(s *config.State, cmd command.Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("missing arguments for login command")
	} else if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	}
	name := cmd.Args[0]
	_, err := s.DB.SelectUser(context.Background(), name)
	if err == nil {
		os.Exit(1)
	}
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		os.Exit(1)
	}
	s.Cfg.SetUser(user.Name)
	return nil
}

func handlerReset(s *config.State, cmd command.Command) error {
	if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	}
	err := s.DB.DeleteAll(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func handlerUsers(s *config.State, cmd command.Command) error {
	if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	}
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.Cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}

func handlerAgg(s *config.State, cmd command.Command) error {
	if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	} else if len(cmd.Args) < 1 {
		return errors.New("not enough arguments")
	}
	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return nil
	}
	fmt.Printf("Collecting feeds every %vs\n", duration.Seconds())
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		fmt.Println("Starting fetch...")
		scrapeFeeds(s)
	}
}

func handlerAddFeed(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) > 2 {
		return errors.New("too many arguments for login command")
	} else if len(cmd.Args) < 2 {
		return errors.New("not enough arguments")
	}
	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerFeeds(s *config.State, cmd command.Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("too many arguments for login command")
	}
	feeds, err := s.DB.SelectAllFeedsWithUsername(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(feeds)
	return nil
}

func handlerFollow(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	} else if len(cmd.Args) < 1 {
		return errors.New("not enough arguments")
	}
	feed, err := s.DB.SelectFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return errors.New("feed was not found")
	}
	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func handlerFollowing(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	}
	feedFollows, err := s.DB.SelectAllFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, item := range feedFollows {
		fmt.Println(item.Name)
	}
	return nil
}

func handlerUnfollow(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return errors.New("too many arguments for login command")
	}
	feed, err := s.DB.SelectFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	err = s.DB.RemoveFeedFollow(context.Background(), database.RemoveFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}
	return nil
}
