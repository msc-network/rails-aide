package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration struct for app
// type Configuration struct {
// 	BaseDir     string
// 	FrontendDir string
// 	Admin       bool
// 	Vue         bool
// 	Rails       bool
// }

// Config file is JSON
type Config struct {
	BaseDir     string `json:"BaseDir"`
	FrontendDir string `json:"FrontendDir"`
	Admin       bool   `json:"Admin"`
	Vue         bool   `json:"Vue"`
	Rails       bool   `json:"Rails"`
}

const defaultFilename string = "config.json"

// load Config file into app
func loadConfig() Config {
	userConfig, err := os.Open("builder.config.json")
	if userConfig != nil {
		jsonParser := json.NewDecoder(userConfig)
		err = jsonParser.Decode(&config)
		check(err)
		return config
	}
	configFile, err := os.Open(defaultFilename)
	defer configFile.Close()
	defer userConfig.Close()
	check(err)
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	check(err)
	return config
}

// Write config file locally
func installConfigFile() {
	file, err := ioutil.ReadFile(defaultFilename)
	check(err)
	fmt.Printf("Installing local config file (config.example.json)\nTo use rename to builder.config.json\nSee docs for configuration details\n")
	writable := []byte(file)
	err2 := ioutil.WriteFile("builder.config.example.json", writable, 0644)
	check(err2)
}
