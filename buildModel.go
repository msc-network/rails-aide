package main

import (
	"bufio"
	"fmt"
	"os"
)

// var buildingModel map[string]interface{}
var modelString []string

// https://golang.org/pkg/bufio/#Scanner

func buildModel() {
	// build user model arguments from user input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please input an attribute:")
		attr, _ := reader.ReadString('\n')
		fmt.Println("Please input an active record type for this attribute:")
		attrType, _ := reader.ReadString('\n')
		for attr == "q" {
			break
		}
		fmt.Printf("you entered: \n" + attr + ":" + attrType)
		// modelString := append(modelString, attr+":"+attrType)
		fmt.Println(modelString)
	}
}
