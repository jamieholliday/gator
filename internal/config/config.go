package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	homeDir, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error getting home directory: %v", err)
	}

	c.CurrentUserName = userName
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error marshalling config: %v", err)
	}

	file, err := os.Create(homeDir)
	if err != nil {
		return fmt.Errorf("Error creating config file: %v", err)
	}

	defer file.Close()
	_, err = file.Write(data)
	return err
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func Read() (Config, error) {
	homeDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Error getting home directory: %v", err)
	}

	jsonConfig, err := os.Open(homeDir + "/.gatorconfig.json")
	if err != nil {
		return Config{}, fmt.Errorf("Error opening config file: %v", err)
	}

	defer jsonConfig.Close()

	byteValue, err := io.ReadAll(jsonConfig)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading config file: %v", err)
	}
	var config Config

	json.Unmarshal(byteValue, &config)
	return config, nil
}
