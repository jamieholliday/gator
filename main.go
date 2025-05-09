package main

import (
	"fmt"
	"github.com/jamieholliday/gator/internal/config"
	"os"
)

func main() {
	configFile, err := config.Read()
	fmt.Printf("Config file: %v\n", configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
		os.Exit(1)
	}
	seterr := configFile.SetUser("jamie")

	if seterr != nil {
		fmt.Fprintf(os.Stderr, "Error setting user: %v\n", seterr)
		os.Exit(1)
	}

	configFile2, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Config file: %v\n", configFile2)

}
