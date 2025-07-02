package main

import (
	"errors"

	"github.com/DXICIDE/Gator/internal/config"
	"github.com/DXICIDE/Gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandHandlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.commandHandlers[cmd.name]; ok {
		return handler(s, cmd)
	}
	return errors.New("no command found")

}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandHandlers[name] = f
}
