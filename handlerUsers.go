package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("GetUsers function problem")
		return err
	}
	for i := 0; i < len(users); i++ {
		if s.cfg.CurrentUserName == users[i].Name {
			fmt.Printf("* %v (current)\n", users[i].Name)
		} else {
			fmt.Printf("* %v\n", users[i].Name)
		}
	}
	return nil
}
