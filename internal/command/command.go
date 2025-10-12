package command

import (
	"os"

	"github.com/hrncacz/go-gator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Commands map[string]func(*config.State, Command) error
}

func (c Commands) Run(state *config.State, cmd Command) {
	f, exist := c.Commands[cmd.Name]
	if !exist {
		os.Exit(1)
	}
	err := f(state, cmd)
	if err != nil {
		os.Exit(1)
	}
	return
}

func (c Commands) Register(name string, f func(*config.State, Command) error) {
	c.Commands[name] = f
}

func Init() Commands {
	return Commands{
		Commands: map[string]func(*config.State, Command) error{},
	}
}
