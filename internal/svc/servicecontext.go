package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"taka-api/internal/config"
	"taka-api/internal/domains/accesscontrol"
)

type ServiceContext struct {
	Config        config.Config
	MySqlConn     sqlx.SqlConn
	AccessCtrlSvc *accesscontrol.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	mySqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	accessCtrlSvc, err := accesscontrol.NewService(
		accesscontrol.WithSsoAuthenticator(),
		accesscontrol.WithMySqlUserRepo(mySqlConn))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:        c,
		MySqlConn:     sqlx.NewMysql(c.Mysql.DataSource),
		AccessCtrlSvc: accessCtrlSvc,
	}
}
