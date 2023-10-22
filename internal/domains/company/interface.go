package company

import "context"

type CompanyRepo interface {
	Register(ctx context.Context, register *Register) (*Register, error)
	GetByCompanyId(ctx context.Context, company *Company) (*Company, error)
	GetByHeadUserId(ctx context.Context, headUser *HeadUser) (*Company, error)
}
