package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AttrNameModel = (*customAttrNameModel)(nil)

type (
	// AttrNameModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAttrNameModel.
	AttrNameModel interface {
		attrNameModel
		withSession(session sqlx.Session) AttrNameModel
	}

	customAttrNameModel struct {
		*defaultAttrNameModel
	}
)

// NewAttrNameModel returns a model for the database table.
func NewAttrNameModel(conn sqlx.SqlConn) AttrNameModel {
	return &customAttrNameModel{
		defaultAttrNameModel: newAttrNameModel(conn),
	}
}

func (m *customAttrNameModel) withSession(session sqlx.Session) AttrNameModel {
	return NewAttrNameModel(sqlx.NewSqlConnFromSession(session))
}
