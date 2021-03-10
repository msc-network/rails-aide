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

// re: filenames, ! = singular model, ? = plural model
const jsonConfig = `
{
  "metadata": {
		"name": "Rails Aide",
		"stub": "rails-aide",
    "basedir": "rails-aide",
    "version": "0.0.2",
    "author": "CromonMS <http://github.com/CromonMS>",
    "description": "A companion for building Rails assets",
    "url": "https://github.com/msc-network/rails-aide",
		"repository": "https://github.com/msc-network/rails-aide",
		"example": "Example: ./rails-aide -m Test -a=false -v=false -u=false -r=true"
	},
	"filenames": {
		"admin": {
			"AdminRecordFile": "!Admin",
			"AdminCollectionFile": "=Admin",
			"AdminNewRecordFile": "New!",
			"AdminEditFile": "Edit!Admin"
		},
		"components": {
			"ComponentFormFile": "!Form",
			"ComponentRecordDetailFile": "!Detail",
			"ComponentListFile": "=List"
		},
		"user": {
			"UserRecordFile": "User!",
			"UserCollectionFile": "User=",
			"UserEditFile": "EditUser!"
		}
	},
  "BaseDir": "/app",
  "FrontendPath": "/javascript/frontend",
  "AdminPagesPath": "/pages/Admin/",
  "UserPagesPath": "/pages/User/",
	"ComponentsPath": "/components/",
	"StorePath": "/store/",
  "Admin": true,
  "Vue": true,
  "Store": true,
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
		Repository  string `json:"repository"`
		Example     string `json:"example"`
	} `json:"metadata"`
	Filenames struct {
		Admin struct {
			AdminRecordFile     string `json:"AdminRecordFile"`
			AdminCollectionFile string `json:"AdminCollectionFile"`
			AdminNewRecordFile  string `json:"AdminNewRecordFile"`
			AdminEditFile       string `json:"AdminEditFile"`
		} `json:"admin"`
		Components struct {
			ComponentFormFile         string `json:"ComponentFormFile"`
			ComponentRecordDetailFile string `json:"ComponentRecordDetailFile"`
			ComponentListFile         string `json:"ComponentListFile"`
		} `json:"components"`
		User struct {
			UserRecordFile     string `json:"UserRecordFile"`
			UserCollectionFile string `json:"UserCollectionFile"`
			UserEditFile       string `json:"UserEditFile"`
		}
	} `json:"filenames"`
	BaseDir        string `json:"BaseDir"`
	FrontendPath   string `json:"FrontendPath"`
	AdminPagesPath string `json:"AdminPagesPath"`
	UserPagesPath  string `json:"UserPagesPath"`
	ComponentsPath string `json:"ComponentsPath"`
	StorePath      string `json:"StorePath"`
	Admin          bool   `json:"Admin"`
	User           bool   `json:"User"`
	Vue            bool   `json:"Vue"`
	Store          bool   `json:"Store"`
	Rails          bool   `json:"Rails"`
}

const defaultCustomFilename string = "ra.config.json"

var jsonConfigBlob = []byte(jsonConfig)

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

// Cleanup old global config file and place in new location
func cleanupConfigFile() {
	checkOldGlobalConfigFolder()

}

// Write local config file
func installLocalConfigFile() {
	err := json.Unmarshal(jsonConfigBlob, &config)
	check(err)
	writableJSON, _ := json.Marshal(config)
	fmt.Printf("Installing local config file (ra.config.json)\nSee docs for configuration details\n")
	writable := []byte(writableJSON)
	err2 := ioutil.WriteFile(defaultCustomFilename, writable, 0754)
	check(err2)
}

// Write a global config file
func installGlobalConfigFile() {
	err := json.Unmarshal(jsonConfigBlob, &config)
	check(err)
	writableJSON, _ := json.Marshal(config)
	fmt.Printf("Installing global config file (ra.config.json) in ~/.rails-go\nSee docs for configuration details\n")
	writable := []byte(writableJSON)
	usr, _ := user.Current()
	appPath := filepath.Join(usr.HomeDir, "/.rails-aide")
	os.Chdir(appPath)
	err2 := ioutil.WriteFile(defaultCustomFilename, writable, 0754)
	check(err2)
}

func checkOldGlobalConfigFolder() {
	usr, _ := user.Current()
	oldConfigPath := filepath.Join(usr.HomeDir, "/.rails-aide")
	// exists, err := os.Stat(oldConfigPath)
	_, err := os.Stat(oldConfigPath)
	if err != nil {
		println("os.Stat(): error for folder name ", oldConfigPath)
		println("and error is : ", err.Error())
		if os.IsNotExist(err) {
			println("Directory Does not exists.")
		}
	}
}

// Cleanup old global config file and place in new location
func cleanupOldGlobalFolder() {
	err := json.Unmarshal(jsonConfigBlob, &config)
	check(err)
	// usr, _ := user.Current()
	// oldConfigPath := filepath.Join(usr.HomeDir, "/.rails-aide")

}

// Check existence of app folder.
func checkOrCreateGlobalConfigFolder() {
	err := json.Unmarshal(jsonConfigBlob, &config)
	check(err)
	usr, _ := user.Current()
	appPath := filepath.Join(usr.HomeDir, "./config/rails-aide")
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		fmt.Printf("Creating global app dir\n\n")
		fmt.Printf("Global config installed in: %s\n\n", appPath)
		err = os.MkdirAll(appPath, 0754)
		installGlobalConfigFile()
	}
}
