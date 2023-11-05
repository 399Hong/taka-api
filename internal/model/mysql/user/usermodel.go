package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		UpdateUserNameAndMobile(ctx context.Context, newData *User) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) UpdateUserNameAndMobile(ctx context.Context, newData *User) error {
	if newData == nil {
		return errors.New("user information is nil")
	}
	if newData.Id == 0 || newData.FirstName == "" || newData.LastName == "" || newData.Mobile == "" {
		return fmt.Errorf("unable to update user info because either id, first last, last name or mobile is missing %v", newData)
	}
	query := fmt.Sprintf("update %s set `first_name` = ?, `last_name` = ?, `mobole` = ? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, newData.FirstName, newData.LastName, newData.Mobile, newData.Id)
	return err
}
