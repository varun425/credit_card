package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	BannedCountries []string `json:"banned_countries"`
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		filePath := "config/banned_countries.json"

		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err.Error())
			return nil
		}


		instance = &Config{}
		err = json.Unmarshal(fileContent, &instance.BannedCountries)
		if err != nil {
			log.Fatalf("Error decoding config file: %s", err.Error())
			return nil
		}
	}

	return instance
}
