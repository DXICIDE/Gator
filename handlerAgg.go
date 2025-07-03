package main

import (
	"context"
	"fmt"

	fetch "github.com/DXICIDE/Gator/internal/fetchFeed"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetch.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
