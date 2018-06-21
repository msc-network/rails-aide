package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

// Config file is JSON
type Config struct {
	Metadata struct {
		Name        string `json:"name"`
		BaseDir     string `json:"basedir"`
		Version     string `json:"version"`
		Author      string `json:"author"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"metadata"`
	BaseDir        string `json:"BaseDir"`
	FrontendPath   string `json:"FrontendPath"`
	AdminPagesPath string `json:"AdminPagesPath"`
	UserPagesPath  string `json:"UserPagesPath"`
	ComponentsPath string `json:"ComponentsPath"`
	Admin          bool   `json:"Admin"`
	Vue            bool   `json:"Vue"`
	Rails          bool   `json:"Rails"`
}

const defaultConfigFile string = "config.json"
const defaultCustomFilename string = "builder.config.json"

// load Config file into app
func loadConfig() Config {
	// Check for global config and use if present.
	globalConfigFile := filepath.Join("~/.rails-aide", defaultCustomFilename)

	// fmt.Println("~/.rails-aide", defaultCustomFilename)

	globalConfig, _ := os.Open(globalConfigFile)
	// check(err)
	defer globalConfig.Close()
	if globalConfig != nil {
		jsonParser := json.NewDecoder(globalConfig)
		err := jsonParser.Decode(&config)
		check(err)
		return config
	}
	// Check for local config and use if present.
	localConfig, _ := os.Open(defaultCustomFilename)
	// check(err)
	defer localConfig.Close()
	if localConfig != nil {
		jsonParser := json.NewDecoder(localConfig)
		err := jsonParser.Decode(&config)
		check(err)
		return config
	}
	// Otherwise use built in config.22
	configFile, err := os.Open(defaultConfigFile)
	check(err)
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	check(err)
	return config
}

// Write local config file
func installLocalConfigFile() {
	file, err := ioutil.ReadFile(defaultConfigFile)
	check(err)
	fmt.Printf("Installing local config file (builder.config.json)\nSee docs for configuration details\n")
	writable := []byte(file)
	err2 := ioutil.WriteFile(defaultCustomFilename, writable, 0754)
	check(err2)
}

// Write a global config file
func installGlobalConfigFile() {
	file, err := ioutil.ReadFile(defaultConfigFile)
	check(err)
	fmt.Printf("Installing global config file (builder.config.json) in ~/.rails-go\nSee docs for configuration details\n")
	writable := []byte(file)
	err2 := ioutil.WriteFile(defaultCustomFilename, writable, 0754)
	check(err2)
}

// Check existence of app folder.
func checkOrCreateGlobalAppFolder() {
	file, err := ioutil.ReadFile(defaultConfigFile)
	check(err)
	usr, err := user.Current()
	appPath := filepath.Join(usr.HomeDir, "/.rails-aide")
	fmt.Printf("App config installed in: %s\n\n", appPath)
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		fmt.Printf("Creating global app dir\n")
		err = os.MkdirAll(appPath, 0754)
		writable := []byte(file)
		os.Chdir(appPath)
		err = ioutil.WriteFile(defaultCustomFilename, writable, 0754)
		check(err)
	}
}
