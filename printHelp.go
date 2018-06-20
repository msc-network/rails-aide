package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func printHelp() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
