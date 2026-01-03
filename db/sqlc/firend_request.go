package db

import (
	"Butterfly/db/models"
	"context"
)

const createFriendRequest = `
INSERT INTO friend_requests (
    from_user_id, to_user_id, request_desc
) VALUES (
    $1, $2, $3
)
`

type CreateFriendRequestParams struct {
	FromUserID  int32  `json:"from_user_id"`
	ToUserID    int32  `json:"to_user_id"`
	RequestDesc string `json:"request_desc"`
}

func (q *Queries) CreateFriendRequest(ctx context.Context, arg *CreateFriendRequestParams) error {
	_, err := q.db.ExecContext(ctx, createFriendRequest, arg.FromUserID, arg.ToUserID, arg.RequestDesc)
	return err
}

const getFriendRequest = `
SELECT EXISTS (
	SELECT 1
	FROM friend_requests
	WHERE from_user_id = $1 AND to_user_id = $2 AND status = 1
)
`

type ExistsFriendRequestParams struct {
	FromUserID int32 `json:"from_user_id"`
	ToUserID   int32 `json:"to_user_id"`
}

// 检查申请是否存在
func (q *Queries) ExistsFriendRequest(ctx context.Context, arg *ExistsFriendRequestParams) (bool, error) {
	row := q.db.QueryRowxContext(ctx, getFriendRequest, arg.FromUserID, arg.ToUserID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (q *Queries) GetFriendRequest(ctx context.Context) (models.FriendRequest, error) {
	// q.db.QueryRowxContext(ctx, "")
	return models.FriendRequest{}, nil
}

const listFriendRequestByPending = `
SELECT id, from_user_id, to_user_id, request_desc, status, created_at, updated_at FROM friend_requests 
WHERE 
	to_user_id = $1 AND status = 1
ORDER BY created_at DESC
`

func (q *Queries) ListFriendRequestByPending(ctx context.Context, toUserID int32) ([]models.FriendRequest, error) {
	rows, err := q.db.QueryxContext(ctx, listFriendRequestByPending, toUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.FriendRequest{}
	for rows.Next() {
		var i models.FriendRequest
		if err := rows.Scan(
			&i.ID,
			&i.FromUserID,
			&i.ToUserID,
			&i.RequestDesc,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
