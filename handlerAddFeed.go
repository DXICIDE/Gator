package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DXICIDE/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
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
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feedParams.UserID = user.ID
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
