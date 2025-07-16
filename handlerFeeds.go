package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, v := range feeds {
		fmt.Println(v.FeedName)
		fmt.Println(v.Url)
		fmt.Println(v.UserName)
	}
	return nil
}
