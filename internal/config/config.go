package config

import (
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	home = home + configFileName
	fmt.Println(home)
	return home, nil
}
