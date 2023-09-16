package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"taka-api/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	MySqlConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MySqlConn: sqlx.NewMysql(c.Mysql.DataSource),
	}
}
