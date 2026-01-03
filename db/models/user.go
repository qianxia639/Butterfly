package models

import "time"

// 用户表
type User struct {
	// 用户ID
	ID int32 `json:"id" db:"id"`
	// 用户名
	Username string `json:"username" db:"username"`
	// 用户昵称
	Nickname string `json:"nickname" db:"nickname"`
	// 用户密码
	Password string `json:"password" db:"password"`
	// 用户邮箱
	Email string `json:"email" db:"email"`
	// 用户性别, 1:男, 2:女, 3: 未知
	Gender int8 `json:"gender" db:"gender"`
	// 出生日期
	Brithday time.Time `json:"brithday" db:"brithday"`
	// 头像URL
	AvatarUrl string `json:"avatar_url" db:"avatar_url"`
	// 个性签名
	Signature string `json:"signature" db:"signature"`
	// 上次密码更新时间
	PasswordChangedAt time.Time `json:"password_changed_at" db:"password_changed_at"`
	// 创建时间
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
