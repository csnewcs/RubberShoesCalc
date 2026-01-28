package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token   string `json:"token"`
	AdminId string `json:"adminId"`
}

func LoadConfig(path string) (Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(file, &config)
	return config, err
}
