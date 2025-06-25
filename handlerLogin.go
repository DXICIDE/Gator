package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 2 {
		err := errors.New("username is required")
		return err
	}
	err := s.cfg.SetUser(cmd.arguments[2])
	if err != nil {
		return err
	}
	fmt.Println("User has been set!")
	return nil
}
