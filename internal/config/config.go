package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hrncacz/go-gator/internal/database"
)

type State struct {
	DB  *database.Queries
	Cfg *Config
}

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (*Config, error) {
	configPath, err := getConfigFilepath()
	if err != nil {
		return nil, errors.New("config file was not found")
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.New("cannot read config file")
	}
	config := &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	if err := write(c); err != nil {
		return fmt.Errorf("issue with writing file: %s", err)
	}
	return nil
}

func getConfigFilepath() (string, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}

	configPath := filepath.Join(userPath, configFileName)
	if _, err = os.Stat(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}

func write(cfg *Config) error {
	configFile, err := getConfigFilepath()
	if err != nil {
		return err
	}
	jsonString, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	if err = os.WriteFile(configFile, []byte(jsonString), 0666); err != nil {
		return err
	}
	return nil
}

func InitState(cfg *Config) *State {
	return &State{Cfg: cfg}
}
