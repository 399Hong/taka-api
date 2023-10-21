package company

import (
	"context"

	"taka-api/internal/svc"
	"taka-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCompanyLogic {
	return &RegisterCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterCompanyLogic) RegisterCompany(req *types.CompanyRegisterationReq) (resp *types.CompanyRegisterationResp, err error) {
	// todo: add your logic here and delete this line

	return &types.CompanyRegisterationResp{}, nil
}
