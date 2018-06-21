package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

var (
	admin               bool
	model               string
	rails               bool
	vue                 bool
	userConfirmation    string
	confirmed           bool
	installLocalConfig  bool
	installGlobalConfig bool
	config              Config
)

func main() {
	flag.Parse()

	// if user does not supply flags, print help
	if flag.NFlag() == 0 {
		printHelp()
	}

	if installGlobalConfig == true {
		installGlobalConfigFile()
		os.Exit(1)
	}

	if installLocalConfig == true {
		installLocalConfigFile()
		os.Exit(1)
	}

	if model == "" {
		fmt.Printf("Model is missing, please re-run\n")
		os.Exit(1)
	}

	fmt.Printf("Creating model scaffold for: %s\n", model)
	fmt.Printf("Admin files?: %t\n", admin)
	fmt.Printf("Vue files?: %t\n", vue)
	fmt.Printf("Run rails commands?: %t\n", rails)

	// Ask to continue
	confirmUserActions("Continue?\n", 3)

	if confirmed == true {
		// If --rails is true then run Rails command too
		if rails == true {
			fmt.Printf("Running Rails commands..\n")
		}

		// If --vue is true then create Vue files
		if vue == true {
			createVueFiles()
		}
	} else {
		fmt.Printf("Cancelled by user\n")
	}
}

func init() {
	checkOrCreateGlobalAppFolder()
	config := loadConfig()

	flag.BoolVarP(&admin, "admin", "a", config.Admin, "Set whether Admin files are created")
	flag.BoolVarP(&installLocalConfig, "config", "c", false, "Install local config file")
	flag.BoolVarP(&installGlobalConfig, "gconfig", "g", false, "Install global config file")
	flag.StringVarP(&model, "model", "m", "", "Specify the name of the Model you'd like to create")
	flag.BoolVarP(&rails, "rails", "r", config.Rails, "Run rails generators")
	flag.BoolVarP(&vue, "vue", "v", config.Vue, "Set whether Vue files are created")
}

func confirmUserActions(s string, tries int) bool {
	reader := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)

		res, err := reader.ReadString('\n')
		check(err)

		if len(res) < 2 {
			continue
		}

		if strings.ToLower(strings.TrimSpace(res))[0] == 'n' {
			confirmed = false
			return confirmed
		}

		if strings.ToLower(strings.TrimSpace(res))[0] == 'y' {
			confirmed = true
			return confirmed
		}
	}

	confirmed = false
	return confirmed
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
