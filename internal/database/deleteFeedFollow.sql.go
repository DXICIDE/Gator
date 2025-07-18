// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: deleteFeedFollow.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE feed_id = $1 AND user_id = $2
`

type DeleteFeedFollowParams struct {
	FeedID uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.FeedID, arg.UserID)
	return err
}
