package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"taka-api/internal/domains/accesscontrol"
	"taka-api/internal/svc"
	"taka-api/internal/types"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// https://developers.google.com/oauthplayground/?code=4%2F0Adeu5BXqxXjO2ibtouj5WP7giY6PZp5Mqu-luxlwT_UshiQiAlP_IR1Ns6gnedH8Y5mb-Q&scope=profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile
// get token need to add email scope
func (l *SignUpLogic) SignUp(req *types.SignUpRequest) (resp *types.SignUpResponse, err error) {
	l.Logger.Debugf("Signing up user %+v", req)
	accessor, err := l.svcCtx.AccessCtrlSvc.SignUp(l.ctx, accesscontrol.AccessMethod{
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

	return &types.SignUpResponse{
		AccessToken:  accessToken,
		RefreshToken: "NotImplemented",
		UserInfo:     nil,
	}, nil
}
