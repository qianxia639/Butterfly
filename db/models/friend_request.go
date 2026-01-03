package models

import "time"

// 好友申请表
type FriendRequest struct {
	// 请求ID
	ID int32 `json:"id" db:"id"`
	// 申请者ID
	FromUserID int32 `json:"from_user_id" db:"from_user_id"`
	// 接收者ID
	ToUserID int32 `json:"to_user_id" db:"to_user_id"`
	// 请求信息
	RequestDesc string `json:"request_desc" db:"request_desc"`
	// 请求状态, 1: 待处理, 2: 已同意, 3: 已拒绝
	Status int8 `json:"status" db:"status"`
	// 请求时间
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// 变更时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
