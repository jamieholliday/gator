package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jamieholliday/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
		os.Exit(1)
	}

	programState := &state{
		cfg: &configFile,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := cliArgs[1]
	cmdArgs := cliArgs[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
