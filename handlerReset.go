package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	s.db.Resetu(context.Background())
	s.db.Resetf(context.Background())
	s.db.Resetfollow(context.Background())
	fmt.Println("Database has been reset!")
	return nil
}
