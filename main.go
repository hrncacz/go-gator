package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hrncacz/go-gator/internal/command"
	"github.com/hrncacz/go-gator/internal/config"
	"github.com/hrncacz/go-gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("chybe: %s\n", err)
		os.Exit(1)
	}
	state := config.InitState(cfg)
	cmd := command.Init()
	cmd.Register("login", handlerLogin)
	cmd.Register("register", handlerRegister)
	cmd.Register("reset", handlerReset)
	cmd.Register("users", handlerUsers)
	cmd.Register("agg", handlerAgg)
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Panic(err)
	}
	dbQueries := database.New(db)
	state.DB = dbQueries
	userCommand := command.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	cmd.Run(state, userCommand)
	fmt.Println(cfg)
	os.Exit(0)
}
