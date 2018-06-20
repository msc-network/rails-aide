package main

import (
	"fmt"
	"io/ioutil"

	inflection "github.com/jinzhu/inflection"
)

var configuration = Configuration{}
var path = configuration.BaseDir

// Create AdminRecordFile - LabelAdmin (admin)
// Create AdminListFile - LabelsAdmin (admin)
// Create AdminNewRecordFile - NewLabel (admin)
// Create AdminEditFile - EditLabelAdmin (admin)
// Create ComponentFormFile - LabelForm (components)
// Create ComponentRecordDetailFile - LabelDetail (components)
// Create ComponentListFile - LabelsList (components)

func createVueFiles() {
	fmt.Printf("Creating Vue Files..\n")
	if admin == true {
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
	err := ioutil.WriteFile(path+model+"Admin.vue", writable, 0644)
	check(err)
}

func createAdminListFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "Admin"))
	err := ioutil.WriteFile(path+inflection.Plural(model)+"Admin.vue", writable, 0644)
	check(err)
}

func createAdminNewRecordFile() {
	writable := []byte(writeTemplate("New" + model))
	err := ioutil.WriteFile(path+"New"+model+".vue", writable, 0644)
	check(err)
}

func createAdminEditFile() {
	writable := []byte(writeTemplate("Edit" + model + "Admin"))
	err := ioutil.WriteFile(path+"Edit"+model+"Admin.vue", writable, 0644)
	check(err)
}

func createComponentFormFile() {
	writable := []byte(writeTemplate(model + "Form"))
	err := ioutil.WriteFile(path+model+"Form.vue", writable, 0644)
	check(err)
}

func createComponentRecordDetailFile() {
	writable := []byte(writeTemplate(model + "Detail"))
	err := ioutil.WriteFile(path+model+"Detail.vue", writable, 0644)
	check(err)
}

func createComponentListFile() {
	writable := []byte(writeTemplate(inflection.Plural(model) + "List"))
	err := ioutil.WriteFile(path+inflection.Plural(model)+"List.vue", writable, 0644)
	check(err)
}
