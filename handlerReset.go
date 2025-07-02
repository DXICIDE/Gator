package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	s.db.Reset(context.Background())
	fmt.Println("Database has been reset!")
	return nil
}
