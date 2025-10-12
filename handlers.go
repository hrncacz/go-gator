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
	"github.com/hrncacz/go-gator/internal/rss"
)

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
	}
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *config.State, cmd command.Command) error {
	if len(cmd.Args) > 2 {
		return errors.New("too many arguments for login command")
	} else if len(cmd.Args) < 2 {
		return errors.New("not enough arguments")
	}
	user, err := s.DB.SelectUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return errors.New("user was not found")
	}

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
