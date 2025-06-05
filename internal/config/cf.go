package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

// to write config into the json file path
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

func (C *Config) SetUser(user_name string) error {
	C.Current_user_name = user_name
	err := write(C)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user home directory: %v", err)
	}
	jsonFilePath := filepath.Join(homeDir, configFileName)
	return jsonFilePath, nil
}

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

const configFileName = ".gatorconfig.json"

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
