package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration struct for app
type Configuration struct {
	BaseDir     string
	FrontendDir string
	Admin       bool
	Vue         bool
	Rails       bool
}

const filename string = "config.json"

func readConfig() string {
	f, err := os.Open(filename)
	check(err)
	readFile := bufio.NewReader(f)
	values, err := readFile.Peek(400)
	check(err)
	// configValues := []byte(readFile)
	// fmt.Print("Config: &s\n", string(values))
	stringValues := string(values)
	return stringValues
}

// load Config file into app
// TODO: Make this work! and use values from config.json before variables.
func loadConfig() {
	// readConfig()
	file, err := os.Open(filename)
	check(err)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	// configValues := []byte()
	check(err)
	// fmt.Print("Config: &s\n", string(decoder))
}

// Config file is JSON
// type Config struct {
// 	BaseDir     string `json:"BaseDir"`
// 	FrontendDir string `json:"FrontendDir"`
// 	Admin       bool   `json:"Admin"`
// 	Vue         bool   `json:"Vue"`
// 	Rails       bool   `json:"Rails"`
// }

// Write config file locally
func installConfigFile() {
	file, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Printf("Installing local config file (config.example.json)\nTo use rename to builder.config.json\nSee docs for configuration details")
	writable := []byte(file)
	err2 := ioutil.WriteFile("builder.config.example.json", writable, 0644)
	check(err2)
}
