package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 2 {
		err := errors.New("username is required")
		return err
	}

	//Checks if the user exists
	_, ok := s.db.GetUser(context.Background(), cmd.arguments[2])
	if ok != nil {
		fmt.Println("User does not exist!")
		os.Exit(1)
	}

	//Setting the user as logged
	err := s.cfg.SetUser(cmd.arguments[2])
	if err != nil {
		return err
	}
	fmt.Println("User has been set!")
	return nil
}
