package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Config struct {
	LocalDirection  string   `json:"local_save_path"`
	Destinations    []string `json:"destinations"`
	HistoryFilePath string   `json:"history_file_path"`
	Protocols       []string `json:"protocols"`
}

var (
	AppConfig Config
	mutex     sync.Mutex
)

func LoadConfig() error {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &AppConfig)
}

func AddDestination(ip string) error {
	mutex.Lock()
	defer mutex.Unlock()
	AppConfig.Destinations = append(AppConfig.Destinations, ip)
	err := SaveConfig()
	if err != nil {
		return err
	}
	return nil
}

func SaveConfig() error {
	data, err := json.MarshalIndent(AppConfig, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("config.json", data, 0644)
}
