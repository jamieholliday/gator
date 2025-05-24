package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jamieholliday/gator/internal/config"
	"github.com/jamieholliday/gator/internal/database"
)

// needs importing for but not you dont use it directly
import _ "github.com/lib/pq"

type state struct {
	cfg *config.Config
	db  *database.Queries
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

	db, err := sql.Open("postgres", programState.cfg.DbUrl)
	dbQueries := database.New(db)
	programState.db = dbQueries

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetAllUsers)

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
