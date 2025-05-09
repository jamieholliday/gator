package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(cfg)
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil
}

func Read() (Config, error) {
	homeDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonConfig, err := os.Open(homeDir)
	if err != nil {
		return Config{}, err
	}

	defer jsonConfig.Close()

	decoder := json.NewDecoder(jsonConfig)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil

}

func write(cfg *Config) error {
	homeDir, err := getConfigFilePath()

	if err != nil {
		return err
	}

	file, err := os.Create(homeDir)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)

	if err != nil {
		return err
	}
	return nil
}
