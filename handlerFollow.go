package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DXICIDE/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	feedParams := database.CreateFeedFollowParams{}
	feedParams.CreatedAt = time.Now()
	feedParams.UpdatedAt = time.Now()
	feedParams.ID = uuid.New()
	feedParams.UserID = user.ID

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[2])
	if err != nil {
		return err
	}
	feedParams.FeedID = feed.ID

	createFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Println(createFeedFollow.FeedName)
	fmt.Println(createFeedFollow.UserName)
	return nil
}
