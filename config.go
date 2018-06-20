package main

import (
	"encoding/json"
	"os"
)

// Configuration struct for app
type Configuration struct {
	BaseDir     string
	FrontendDir string
}

func loadConfig() {
	filename = "config.json"
	file, err := os.Open(filename)
	check(err)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	check(err)
}
