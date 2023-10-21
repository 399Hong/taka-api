package company

import "context"

type Repo interface {
	Register(ctx context.Context, accessor *Register) (*Register, error)
	GetByCompanyId(ctx context.Context, company *Company) (*Company, error)
	GetByHeadUserId(ctx context.Context, headUser *HeadUser) (*Company, error)
}
