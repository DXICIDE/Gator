package main

import (
	"context"
	"fmt"

	"github.com/DXICIDE/Gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	followData, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(followData) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	for _, v := range followData {
		fmt.Println(v.Name)
		fmt.Println(v.UserName)
	}
	return nil
}
