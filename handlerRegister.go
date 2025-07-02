package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/DXICIDE/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 2 {
		err := errors.New("username is required")
		return err
	}

	//sets the user database params
	user := database.CreateUserParams{}
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Name = cmd.arguments[2]

	_, ok := s.db.GetUser(context.Background(), cmd.arguments[2])
	if ok == nil {
		fmt.Println("User with this name already exists!")
		os.Exit(1)
	}

	//creates user with empty context arg and set user
	s.db.CreateUser(context.Background(), user)
	err := s.cfg.SetUser(cmd.arguments[2])
	if err != nil {
		fmt.Println("SetUset problem")
		os.Exit(1)
	}
	fmt.Println("User has been created!")
	return nil
}
