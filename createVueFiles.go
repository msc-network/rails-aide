package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/jinzhu/inflection"
)

var frontendPath string
var adminPath string
var componentsPath string
var userPath string
var templateName string

func createVueFiles() {
	// Checks for existence of frontend directory and if not, creates.
	frontendPath = filepath.Join("./", config.BaseDir, config.FrontendPath)
	if _, err := os.Stat(frontendPath); os.IsNotExist(err) {
		fmt.Printf("Creating frontend directory structure\n")
		os.MkdirAll(frontendPath, 0754)
	}

	fmt.Printf("Creating Vue Files..\n")
	if admin == true {
		adminPath = filepath.Join(frontendPath, config.AdminPagesPath, inflection.Plural(model))
		if _, err := os.Stat(adminPath); os.IsNotExist(err) {
			fmt.Printf("Creating admin directory structure\n")
			os.MkdirAll(adminPath, 0754)
		}
		fmt.Printf("Writing admin files..\n")
		createAdminSet()
	}

	componentsPath = filepath.Join(frontendPath, config.ComponentsPath, inflection.Plural(model))
	if _, err := os.Stat(componentsPath); os.IsNotExist(err) {
		fmt.Printf("Creating directory structure\n")
		os.MkdirAll(componentsPath, 0754)
	}
	fmt.Printf("Writing component files..\n")
	createComponentSet()

	if users == true {
		userPath = filepath.Join(frontendPath, config.UserPagesPath, inflection.Plural(model))
		if _, err := os.Stat(userPath); os.IsNotExist(err) {
			fmt.Printf("Creating directory structure\n")
			os.MkdirAll(userPath, 0754)
		}
		fmt.Printf("Writing user files..\n")
		createUserSet()
	}
}

func createAdminSet() {
	singularMatch := regexp.MustCompile("(!)")
	pluralMatch := regexp.MustCompile("(=)")
	adminSet := config.Filenames.Admin
	v := reflect.ValueOf(adminSet)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		filenamestring := fmt.Sprint(values[i])
		if singularMatch.MatchString(filenamestring) {
			templateName = singularMatch.ReplaceAllLiteralString(filenamestring, model)
		} else if pluralMatch.MatchString(filenamestring) {
			templateName = pluralMatch.ReplaceAllLiteralString(filenamestring, inflection.Plural(model))
		}
		fileFullPath := adminPath + "/" + templateName + ".vue"
		createFile(templateName, fileFullPath)
	}
}

func createComponentSet() {
	singularMatch := regexp.MustCompile("(!)")
	pluralMatch := regexp.MustCompile("(=)")
	componentSet := config.Filenames.Components
	v := reflect.ValueOf(componentSet)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		filenamestring := fmt.Sprint(values[i])
		if singularMatch.MatchString(filenamestring) {
			templateName = singularMatch.ReplaceAllLiteralString(filenamestring, model)
		} else if pluralMatch.MatchString(filenamestring) {
			templateName = pluralMatch.ReplaceAllLiteralString(filenamestring, inflection.Plural(model))
		}
		fileFullPath := componentsPath + "/" + templateName + ".vue"
		createFile(templateName, fileFullPath)
	}
}

func createUserSet() {
	singularMatch := regexp.MustCompile("(!)")
	pluralMatch := regexp.MustCompile("(=)")
	userSet := config.Filenames.User
	v := reflect.ValueOf(userSet)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		filenamestring := fmt.Sprint(values[i])
		if singularMatch.MatchString(filenamestring) {
			templateName = singularMatch.ReplaceAllLiteralString(filenamestring, model)
		} else if pluralMatch.MatchString(filenamestring) {
			templateName = pluralMatch.ReplaceAllLiteralString(filenamestring, inflection.Plural(model))
		}
		fileFullPath := userPath + "/" + templateName + ".vue"
		createFile(templateName, fileFullPath)
	}
}

// TODO: [unused] Test generic createFile and make it handle arrays
func createFile(templateName string, fileFullPath string) {
	writable := []byte(writeTemplate(templateName))
	err := ioutil.WriteFile(fileFullPath, writable, 0754)
	check(err)
}
