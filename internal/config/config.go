package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// configFileName specifies the default configuration file name.
const configFileName = ".gatorconfig.json"

// Config holds application-wide configuration settings.
type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

// getConfigFilePath returns the full path to the application's configuration file.
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user home directory: %v", err)
	}
	jsonFilePath := filepath.Join(homeDir, configFileName)
	return jsonFilePath, nil
}

// write persists the given Config struct to the configuration file.
func write(cfg *Config) error {
	raw, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error when converting to raw text from json: %v", err)
	}
	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error trying to write file: %v", err)
	}
	err = os.WriteFile(jsonFilePath, raw, 0644)
	if err != nil {
		return fmt.Errorf("error trying to write json data: %v", err)
	}
	return nil
}

// SetUser updates the current user name in the Config and persists the changes.
func (C *Config) SetUser(user_name string) error {
	C.Current_user_name = user_name
	err := write(C)
	if err != nil {
		return err
	}
	return nil
}

// Read loads and unmarshals the application configuration from the file system.
func Read() (*Config, error) {
	config := &Config{}

	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return config, fmt.Errorf("error reading jsonpath: %v", err)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		fmt.Println(err)
	}

	return config, nil
}

// ReadJsonConfigFile reads and prints the raw content of the configuration file.
func ReadJsonConfigFile() error {
	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("error reading jsonpath: %v", err)
	}
	fmt.Println(string(data))
	return nil
}
