package company

import "context"

type CompanyRepo interface {
	Register(ctx context.Context, register *Register) (*Register, error)
	GetByCompanyId(ctx context.Context, companyId int64) (*Company, error)
	GetByHeadUserId(ctx context.Context, headUserID int64) (*Company, error)
	UpdateCompanyHeadUserInfo(ctx context.Context, headUser *HeadUser) error
}

type CompanyUserRepo interface {
	UpdateCompanyHeadUserInfo(ctx context.Context, headUser *HeadUser) error
}
