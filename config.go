package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const jsonConfig = `
{
  "metadata": {
		"name": "Rails Aide",
		"stub": "rails-aide",
    "basedir": "/.rails-aide",
    "version": "0.0.1",
    "author": "CromonMS <http://github.com/CromonMS>",
    "description": "A companion for building Rails assets",
    "url": "https://github.com/msc-network/rails-aide - TODO: change name"
  },
  "BaseDir": "/app",
  "FrontendPath": "/javascript/frontend",
  "AdminPagesPath": "/pages/Admin/",
  "UserPagesPath": "/pages/User/",
  "ComponentsPath": "/components/",
  "Admin": true,
  "Vue": true,
  "Rails": false
}`

// Config file is JSON
type Config struct {
	Metadata struct {
		Name        string `json:"name"`
		Stub        string `json:"stub"`
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
	usr, _ := user.Current()
	// Check for local config and use if present.
	localConfig, _ := os.Open(defaultCustomFilename)
	defer localConfig.Close()
	if localConfig != nil {
		jsonParser := json.NewDecoder(localConfig)
		err := jsonParser.Decode(&config)
		check(err)
		return config
	}
	// Check for global config and use if present.
	globalConfigFile := filepath.Join(usr.HomeDir, ".rails-aide", defaultCustomFilename)
	globalConfig, _ := os.Open(globalConfigFile)
	defer globalConfig.Close()
	if globalConfig != nil {
		jsonParser := json.NewDecoder(globalConfig)
		err := jsonParser.Decode(&config)
		check(err)
		return config
	}
	// Otherwise use built in config.
	jsonParser := json.NewDecoder(strings.NewReader(jsonConfig))
	err := jsonParser.Decode(&config)
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
	usr, _ := user.Current()
	appPath := filepath.Join(usr.HomeDir, "/.rails-aide")
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		fmt.Printf("Creating global app dir\n\n")
		fmt.Printf("Global config installed in: %s\n\n", appPath)
		err = os.MkdirAll(appPath, 0754)
		writable := []byte(file)
		os.Chdir(appPath)
		err = ioutil.WriteFile(defaultCustomFilename, writable, 0754)
		check(err)
	}
}
