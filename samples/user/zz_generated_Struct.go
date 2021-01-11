// Code generated by go-gendb. DO NOT EDIT.
// go-gendb version: 0.1.0
// source: samples/user/oper.go
package user

import (
	time "time"
)

// UserDetail is an auto-generated return type for method GetDetail
type UserDetail struct {
	// UserId 用户id
	UserId int64 `table:"user_detail" field:"user_id"`
	// Score 用户得分
	Score int32 `table:"user_detail" field:"score"`
	// Balance 用户余额
	Balance int32 `table:"user_detail" field:"balance"`
	// Text 用户详细文本描述
	Text string `table:"user_detail" field:"text"`
}

// User is an auto-generated return type for method GetById
type User struct {
	// Id 用户id
	Id int64 `table:"user" field:"id"`
	// Name 用户名
	Name string `table:"user" field:"name"`
	// Email 用户邮箱
	Email string `table:"user" field:"email"`
	// Phone 用户电话
	Phone string `table:"user" field:"phone"`
	// Age 用户年龄
	Age int32 `table:"user" field:"age"`
	// CreateTime 创建时间
	CreateTime time.Time `table:"user" field:"create_time"`
	// Password 用户密码
	Password string `table:"user" field:"password"`
	// IsAdmin 是否是管理员 0-不是  1-是
	IsAdmin int32 `table:"user" field:"is_admin"`
}
