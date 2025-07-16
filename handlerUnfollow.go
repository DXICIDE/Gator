package main

import (
	"context"
	"fmt"

	"github.com/DXICIDE/Gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[2])
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	userFound, err := s.db.GetUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("couldn't find the user: %w", err)
	}

	params := database.DeleteFeedFollowParams{}
	params.FeedID = feed.ID
	params.UserID = userFound.ID
	s.db.DeleteFeedFollow(context.Background(), params)
	return nil
}
