package user

import (
	"context"
	"taka-api/internal/domains/accesscontrol"

	"taka-api/internal/svc"
	"taka-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	l.Logger.Debugf("loging user %+v", req)
	accessor, err := l.svcCtx.AccessCtrlSvc.LogIn(l.ctx, accesscontrol.AccessMethod{
		AccessOption: accesscontrol.AccessOption(req.AuthType),
		Token:        req.Token,
		Email:        req.Email,
		Password:     req.Password,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	uid, err := accessor.GetId()
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	auth := l.svcCtx.Config.Auth
	accessToken, err := l.svcCtx.AccessCtrlSvc.CreateJwtToken(
		auth.AccessSecret,
		auth.ExpiresInSec,
		uid)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: "NotImplemented",
		UserInfo:     nil,
	}, nil
}
