package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	LocalDirection string   `json:"local_save_path"`
	Destinations   []string `json:"destinations"`
}

var AppConfig Config

func LoadConfig() error {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &AppConfig)
}

func SaveConfig() error {
	data, err := json.MarshalIndent(AppConfig, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("config.json", data, 0644)
}
