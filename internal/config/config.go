package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DB_URL    string `json:"db_url"`
	JwtSecret string `json:"jwt_secret"`
}

const configFileName = ".markableconfig.json"
const db_url = "postgres://mark:pwd@localhost:5433/hospital?sslmode=disable"
const key = "secretKey"

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not determine the home directory: %q", err)
	}
	completePath := filepath.Join(homeDir, configFileName)
	return completePath, nil
}

func Read() (Config, error) {
	configpath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(configpath)
	if err != nil {
		return Config{}, fmt.Errorf("Could not open file: %q", err)
	}
	defer file.Close()
	c := Config{}

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		return Config{}, fmt.Errorf("Error decoding file: %q", err)
	}
	return c, nil
}

func SetConfig() error {
	cfg := &Config{}
	cfg.DB_URL = db_url
	cfg.JwtSecret = key
	return write(cfg)
}

func write(c *Config) error {
	fullpath, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("Error reading file path: %q", err)
	}
	file, err := os.Create(fullpath)
	if err != nil {
		return fmt.Errorf("Error creating file: %q", err)
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(&c)
	if err != nil {
		return fmt.Errorf("Error encoding file: %q", err)
	}
	return nil

}
