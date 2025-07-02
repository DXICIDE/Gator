package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/DXICIDE/Gator/internal/config"
	"github.com/DXICIDE/Gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		err := errors.New("not enough arguments")
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	//making state struct
	stateStuct := new(state)
	stateStuct.cfg = &c

	//creating a connection to database
	db, err := sql.Open("postgres", stateStuct.cfg.DbUrl)
	if err != nil {
		fmt.Println("something went wrong with sqp.Open")
		os.Exit(1)
	}

	//creating *database.Queries and storing it in state
	dbQueries := database.New(db)
	stateStuct.db = dbQueries

	//handling of commands
	commands := new(commands)
	commands.commandHandlers = make(map[string]func(*state, command) error)
	args := new(command)
	args.name = os.Args[1]
	args.arguments = os.Args

	//registering a command
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	err = commands.run(stateStuct, *args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
