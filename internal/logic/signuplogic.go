package logic

import (
	"context"
	"database/sql"
	"errors"
	"taka-api/internal/model/mysql/user"
	"taka-api/internal/svc"
	"taka-api/internal/types"
	"taka-api/pkg/signon"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
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

// // https://developers.google.com/oauthplayground/?code=4%2F0Adeu5BXqxXjO2ibtouj5WP7giY6PZp5Mqu-luxlwT_UshiQiAlP_IR1Ns6gnedH8Y5mb-Q&scope=profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile
// // get token need to add email scope
func (l *SignUpLogic) SignUp(req *types.SignUpRequest) (resp *types.SignUpResponse, err error) {
	resp = new(types.SignUpResponse)
	req.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjgzOGMwNmM2MjA0NmMyZDk0OGFmZmUxMzdkZDUzMTAxMjlmNGQ1ZDEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI0MDc0MDg3MTgxOTIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0MDc0MDg3MTgxOTIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDQ2NTc1NDM5Mjc2NTc3OTg2MTEiLCJoZCI6ImF1Y2tsYW5kdW5pLmFjLm56IiwiZW1haWwiOiJoemhhNzg0QGF1Y2tsYW5kdW5pLmFjLm56IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJWY2J1TXVyWEhNa3ZqT05Sa2FLSjZnIiwiaWF0IjoxNjk0MjQ0NDUyLCJleHAiOjE2OTQyNDgwNTJ9.B23zhpubnTBBxWfPrXjc6KedbMqTveu7jTh5kMRAtYPjZ5XrXYADEOJN5ykvlK0ANipbXHwbOgxcpGzJKwLSWqvnJmOE-yHWxeyJvbsQDn-PTldtDQBSjf8GOcNGmtUTiZj9_ZDBjsZxULmaaaWakEXYiH4QCnY2zPELCwWEZeDRe4ZI5cjl21xvsdAJUfipc_j_7MP2y71wlS5TjeMnEDSX2t1L8SgqohMBM0jCYGE-3Qhaf7WhJWSXY59uBz2iz5k4jXh24vME6DSSxIUa-H1ZdZ8VsEdNzJBRbXKaZa96iyA0j83f3UaGHYvvflYNmpXA0rn55tkJ7dIR4uTgpA"
	auth, err := signon.NewFactory(signon.Google)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	claim, err := auth.Authenticate(l.ctx, req.Token)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	userRepo := user.NewUserModel(l.svcCtx.MySqlConn)

	// extract method
	retrievedUser, err := userRepo.FindOneByEmail(l.ctx, claim.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			l.Logger.Error(err)
			return nil, err
		}
	} else {
		resp.UserInfo = retrievedUser
		return resp, errors.New("user is already registered")
	}

	newUser := &user.User{
		FirstName: "Hongyue",
		LastName:  "Zhang",
		Password:  "1234",
		Mobile:    "00000",
		Email:     claim.Email,
	}
	res, err := userRepo.Insert(l.ctx, newUser)
	if err != nil {
		l.Logger.Error("unable to create user", err)
		return nil, errors.New("unable to create user")
	}
	userId, err := res.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	userInfo, err := userRepo.FindOne(l.ctx, userId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	resp.UserInfo = userInfo
	return resp, nil
}
