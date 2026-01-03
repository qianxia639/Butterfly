package models

import "time"

// 好友关系表
type Friendship struct {
	// 用户ID
	UserID int32 `json:"user_id" db:"user_id"`
	// 好友ID
	FriendID int32 `json:"friend_id" db:"friend_id"`
	// 好友备注
	Remark string `json:"remark" db:"remark"`
	// 创建时间
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
