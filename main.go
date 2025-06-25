package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/DXICIDE/Gator/internal/config"
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

	s := new(state)
	s.cfg = &c
	commands := new(commands)
	commands.commandHandlers = make(map[string]func(*state, command) error)
	args := new(command)
	args.name = os.Args[1]
	args.arguments = os.Args

	commands.register("login", handlerLogin)
	err = commands.run(s, *args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
