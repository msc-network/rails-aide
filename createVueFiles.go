package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	inflection "github.com/jinzhu/inflection"
)

var vuePath string
var adminPath string
var componentsPath string
var userPath string

// Create AdminRecordFile - LabelAdmin (admin)
// Create AdminCollectionFile - LabelsAdmin (admin)
// Create AdminNewRecordFile - NewLabel (admin)
// Create AdminEditFile - EditLabelAdmin (admin)

// Create ComponentFormFile - LabelForm (components)
// Create ComponentRecordDetailFile - LabelDetail (components)
// Create ComponentListFile - LabelsList (components)

// Create UserRecordFile - UserLabel (user)
// Create UserCollectionFile - UserLabels (user)
// Create UserEditFile - EditUserLabel (admin)

func createVueFiles() {
	// Checks for existence of directory and if not, creates.
	vuePath = filepath.Join("./", config.BaseDir, config.FrontendPath)
	if _, err := os.Stat(vuePath); os.IsNotExist(err) {
		fmt.Printf("Creating directory structure\n")
		os.MkdirAll(vuePath, 0754)
	}

	fmt.Printf("Creating Vue Files..\n")
	if admin == true {
		adminPath = filepath.Join(vuePath, config.AdminPagesPath, inflection.Plural(model))
		if _, err := os.Stat(adminPath); os.IsNotExist(err) {
			fmt.Printf("Creating directory structure\n")
			os.MkdirAll(adminPath, 0754)
		}
		fmt.Printf("Writing admin files..\n")
		createAdminRecordFile()
		createAdminCollectionFile()
		createAdminNewRecordFile()
		createAdminEditFile()
	}
	fmt.Printf("Writing component files..\n")
	componentsPath = filepath.Join(vuePath, config.ComponentsPath, inflection.Plural(model))
	if _, err := os.Stat(componentsPath); os.IsNotExist(err) {
		fmt.Printf("Creating directory structure\n")
		os.MkdirAll(componentsPath, 0754)
	}
	createComponentFormFile()
	createComponentRecordDetailFile()
	createComponentListFile()

	fmt.Printf("Writing user files..\n")
	userPath = filepath.Join(vuePath, config.UserPagesPath, inflection.Plural(model))
	if _, err := os.Stat(userPath); os.IsNotExist(err) {
		fmt.Printf("Creating directory structure\n")
		os.MkdirAll(userPath, 0754)
	}
	createUserRecordFile()
	createUserCollectionFile()
	createUserEditFile()
}

// TODO: [unused] Test generic createFile and make it handle arrays
func createFile(templateName string, fileFullPath string) {
	writable := []byte(writeTemplate(templateName))
	err := ioutil.WriteFile(fileFullPath, writable, 0754)
	check(err)
}

// Admin
func createAdminRecordFile() {
	writable := []byte(writeTemplate(model + "Admin.vue"))
	err := ioutil.WriteFile(adminPath+"/"+model+"Admin.vue", writable, 0754)
	check(err)
}

func createAdminCollectionFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "Admin"))
	err := ioutil.WriteFile(adminPath+"/"+inflection.Plural(model)+"Admin.vue", writable, 0754)
	check(err)
}

func createAdminNewRecordFile() {
	writable := []byte(writeTemplate("New" + model))
	err := ioutil.WriteFile(adminPath+"/"+"New"+model+".vue", writable, 0754)
	check(err)
}

func createAdminEditFile() {
	writable := []byte(writeTemplate("Edit" + model + "Admin"))
	err := ioutil.WriteFile(adminPath+"/"+"Edit"+model+"Admin.vue", writable, 0754)
	check(err)
}

// Components
func createComponentFormFile() {
	writable := []byte(writeTemplate(model + "Form"))
	err := ioutil.WriteFile(componentsPath+"/"+model+"Form.vue", writable, 0754)
	check(err)
}

func createComponentRecordDetailFile() {
	writable := []byte(writeTemplate(model + "Detail"))
	err := ioutil.WriteFile(componentsPath+"/"+model+"Detail.vue", writable, 0754)
	check(err)
}

func createComponentListFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "List"))
	err := ioutil.WriteFile(componentsPath+"/"+inflection.Plural(model)+"List.vue", writable, 0754)
	check(err)
}

// User
func createUserRecordFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "List"))
	err := ioutil.WriteFile(userPath+"/"+"User"+model+".vue", writable, 0754)
	check(err)
}

func createUserCollectionFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "List"))
	err := ioutil.WriteFile(userPath+"/"+"User"+inflection.Plural(model)+".vue", writable, 0754)
	check(err)
}

func createUserEditFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "List"))
	err := ioutil.WriteFile(userPath+"/"+"EditUser"+model+".vue", writable, 0754)
	check(err)
}
