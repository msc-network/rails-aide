package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	inflection "github.com/jinzhu/inflection"
)

func init() {

}

// Create AdminRecordFile - LabelAdmin (admin)
// Create AdminListFile - LabelsAdmin (admin)
// Create AdminNewRecordFile - NewLabel (admin)
// Create AdminEditFile - EditLabelAdmin (admin)
// Create ComponentFormFile - LabelForm (components)
// Create ComponentRecordDetailFile - LabelDetail (components)
// Create ComponentListFile - LabelsList (components)

func createVueFiles() {
	// Checks for existence of directory and if not, creates.
	vuePath := filepath.Join("", config.BaseDir, config.FrontendPath)
	if _, err := os.Stat(vuePath); os.IsNotExist(err) {
		fmt.Printf("Creating directory structure\n")
		os.MkdirAll("./"+vuePath, 0754)
	}

	fmt.Printf("Creating Vue Files..\n")
	if admin == true {
		os.Chdir(vuePath)
		fmt.Printf("Writing admin files..\n")
		createAdminRecordFile()
		createAdminListFile()
		createAdminNewRecordFile()
		createAdminEditFile()
	}
	fmt.Printf("Writing component files..\n")
	createComponentFormFile()
	createComponentRecordDetailFile()
	createComponentListFile()

}

func createAdminRecordFile() {
	writable := []byte(writeTemplate(model + "Admin.vue"))
	err := ioutil.WriteFile(model+"Admin.vue", writable, 0754)
	check(err)
}

func createAdminListFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "Admin"))
	err := ioutil.WriteFile(inflection.Plural(model)+"Admin.vue", writable, 0754)
	check(err)
}

func createAdminNewRecordFile() {
	writable := []byte(writeTemplate("New" + model))
	err := ioutil.WriteFile("New"+model+".vue", writable, 0754)
	check(err)
}

func createAdminEditFile() {
	writable := []byte(writeTemplate("Edit" + model + "Admin"))
	err := ioutil.WriteFile("Edit"+model+"Admin.vue", writable, 0754)
	check(err)
}

func createComponentFormFile() {
	writable := []byte(writeTemplate(model + "Form"))
	err := ioutil.WriteFile(model+"Form.vue", writable, 0754)
	check(err)
}

func createComponentRecordDetailFile() {
	writable := []byte(writeTemplate(model + "Detail"))
	err := ioutil.WriteFile(model+"Detail.vue", writable, 0754)
	check(err)
}

func createComponentListFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "List"))
	err := ioutil.WriteFile(inflection.Plural(model)+"List.vue", writable, 0754)
	check(err)
}
