package genModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HomestayOrderModel = (*customHomestayOrderModel)(nil)

type (
	// HomestayOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayOrderModel.
	HomestayOrderModel interface {
		homestayOrderModel
	}

	customHomestayOrderModel struct {
		*defaultHomestayOrderModel
	}
)

// NewHomestayOrderModel returns a model for the database table.
func NewHomestayOrderModel(conn sqlx.SqlConn) HomestayOrderModel {
	return &customHomestayOrderModel{
		defaultHomestayOrderModel: newHomestayOrderModel(conn),
	}
}
