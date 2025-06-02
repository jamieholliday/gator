package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jamieholliday/gator/internal/config"
	"github.com/jamieholliday/gator/internal/database"

	// needs importing for but not you dont use it directly
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			fmt.Errorf("Error getting user by name %s: %v\n", s.cfg.CurrentUserName, err)
		}
		return handler(s, cmd, user)
	}
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
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetAllFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerGetFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handlerDeleteFeedFollows))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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
