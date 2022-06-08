package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel { // 通过接口调用内部model 接口
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}
