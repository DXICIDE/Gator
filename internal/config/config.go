package config

import (
	"encoding/json"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(current_user string) error {
	c.CurrentUserName = current_user
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

func write(c *Config) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	dir, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	home, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(home)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	home = home + "/" + configFileName
	return home, nil
}
