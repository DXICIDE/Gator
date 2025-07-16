package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DXICIDE/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 4 {
		fmt.Println("missing arguments name or url or both")
		os.Exit(1)
	}

	feedParams := database.CreateFeedParams{}
	feedParams.ID = uuid.New()
	feedParams.CreatedAt = time.Now()
	feedParams.UpdatedAt = time.Now()
	feedParams.Name = cmd.arguments[2]
	feedParams.Url = cmd.arguments[3]

	feedParams.UserID = user.ID
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	cmd.arguments[2] = cmd.arguments[3]
	handlerFollow(s, cmd, user)

	fmt.Printf("%+v\n", feed)
	return nil
}
