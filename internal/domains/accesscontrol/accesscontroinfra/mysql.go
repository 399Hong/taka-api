package accesscontroinfra

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"taka-api/internal/domains/accesscontrol"
	"taka-api/internal/model/mysql/user"
)

type MySqlRepo struct {
	user.UserModel
}

func NewMySqlRepo(conn sqlx.SqlConn) *MySqlRepo {
	return &MySqlRepo{user.NewUserModel(conn)}
}

func (m *MySqlRepo) Add(ctx context.Context, accessor *accesscontrol.Accessor) (*accesscontrol.Accessor, error) {
	newUser := &user.User{
		Password: accessor.Password,
		Email:    accessor.Email,
	}
	user, err := m.Insert(ctx, newUser)
	if err != nil {
		return accessor, err
	}
	uid, err := user.LastInsertId()
	if err != nil {
		return accessor, err
	}

	accessor.SetId(uid)
	return accessor, nil
}

func (m *MySqlRepo) GetById(ctx context.Context, id int64) (*accesscontrol.User, error) {
	user, err := m.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return &accesscontrol.User{Id: user.Id, Password: user.Password, Email: user.Email}, nil
}

func (m *MySqlRepo) GetByEmail(ctx context.Context, email string) (*accesscontrol.User, error) {
	user, err := m.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &accesscontrol.User{Id: user.Id, Password: user.Password, Email: user.Email}, nil
}
