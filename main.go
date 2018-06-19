package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

var (
	admin            bool
	model            string
	rails            bool
	vue              bool
	userConfirmation string
	confirmed        bool
)

func main() {
	flag.Parse()

	// if user does not supply flags, print help
	if flag.NFlag() == 0 {
		printHelp()
	}

	fmt.Printf("Creating model scaffold for: %s\n", model)
	fmt.Printf("Admin files?: %t\n", admin)
	fmt.Printf("Vue files?: %t\n", vue)
	fmt.Printf("Run rails commands?: %t\n", rails)

	// Ask to continue
	// fmt.Printf("Continue?\n")
	confirmUserActions("Continue?\n", 10)

	if confirmed == true {
		// If --rails is true then run Rails command too
		fmt.Printf("Running Rails commands..\n")

		// If --vue is true then create Vue files
		if vue == true {
			fmt.Printf("Creating Vue Files..\n")
			createVueFiles()
		}

		// If --admin is true then create Admin files

	} else {
		fmt.Printf("Cancelled by user\n")
	}

}

func init() {
	flag.BoolVarP(&admin, "admin", "a", true, "Set whether Admin files are created")
	flag.StringVarP(&model, "model", "m", "", "Specify the name of the Model you'd like to create")
	flag.BoolVarP(&rails, "rails", "r", false, "Run rails generators")
	flag.BoolVarP(&vue, "vue", "v", true, "Set whether Vue files are created")
}

func confirmUserActions(s string, tries int) bool {

	reader := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)

		res, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if len(res) < 2 {
			continue
		}
		if strings.ToLower(strings.TrimSpace(res))[0] == 'y' {
			confirmed = true
			return confirmed
		}
	}

	confirmed = false
	return confirmed
}

func createVueFiles() {
	fmt.Printf(vueTemplate)
}

func printHelp() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
