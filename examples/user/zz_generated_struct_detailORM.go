// Code generated by go-gendb. DO NOT EDIT.
// go-gendb version: 0.0.1
// source: examples/user/user.db.go
package user

// detailORM
type detailORM struct {
	// UserId 用户id
	UserId int64 `table:"user_detail" field:"user_id"`
	// Score 用户得分
	Score int32 `table:"user_detail" field:"score"`
	// Balance 用户余额
	Balance int32 `table:"user_detail" field:"balance"`
	// Text 用户详细文本描述
	Text string `table:"user_detail" field:"text"`
}

