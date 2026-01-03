package db

import "context"

const existsFriendship = `
SELECT EXISTS(
    SELECT 1 FROM friendships
    WHERE (user_id = $1 AND friend_id = $2)
)
`

type ExistsFriendshipParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
}

// 是否已是好友关系
func (q *Queries) ExistsFriendship(ctx context.Context, arg *ExistsFriendshipParams) (bool, error) {
	row := q.db.QueryRowxContext(ctx, existsFriendship, arg.UserID, arg.FriendID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
