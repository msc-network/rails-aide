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
	cleanupConfig       bool
	config              Config
	confirmed           bool
	installGlobalConfig bool
	installLocalConfig  bool
	model               string
	rails               bool
	store               bool
	userConfirmation    string
	users               bool
	vue                 bool
)

func main() {
	// createAdminFiles()

	fmt.Println(config.Metadata.Name + " | version: " + config.Metadata.Version)
	fmt.Println(config.Metadata.Description + "\n")
	fmt.Println(config.Metadata.URL + "\n")
	fmt.Println(config.Metadata.Example + "\n")

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

	if vue == false {
		if admin == true || users == true {
			fmt.Println("Cannot run admin or user tasks without enabling vue tasks.")
			os.Exit(1)
		}
		fmt.Println("No vue tasks enabled, all other flags will be ignored at this time.")
		os.Exit(1)
	}

	fmt.Printf("Creating model scaffold for: %s\n", model)
	fmt.Printf("Admin files?: %t\n", admin)
	fmt.Printf("Vue files?: %t\n", vue)
	fmt.Printf("Run rails commands?: %t\n", rails)

	// Ask to continue, amount of tries
	confirmUserActions("Continue?\n", 3)

	if confirmed == true {
		// If --rails is true run model builder
		if rails == true {
			buildModel()
		}

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
	checkOrCreateGlobalConfigFolder()
	config := loadConfig()

	flag.BoolVarP(&admin, "admin", "a", config.Admin, "Set whether Admin files are created")
	flag.BoolVarP(&installLocalConfig, "config", "c", false, "Install local config file")
	flag.BoolVarP(&installLocalConfig, "cleanup-config", "d", false, "Clean up config (if you are using an old version)")
	flag.BoolVarP(&installGlobalConfig, "gconfig", "g", false, "Install global config file")
	flag.StringVarP(&model, "model", "m", "", "Specify the name of the Model you'd like to create")
	flag.BoolVarP(&rails, "rails", "r", config.Rails, "Run rails generators")
	flag.BoolVarP(&vue, "vue", "v", config.Vue, "Set whether Vue files are created")
	flag.BoolVarP(&store, "store", "s", config.Store, "Set whether Vue store files are created")
	flag.BoolVarP(&users, "users", "u", config.Vue, "Set whether User files are created")
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
