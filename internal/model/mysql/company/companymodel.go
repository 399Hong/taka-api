package company

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ Model = (*customCompanyModel)(nil)

type (
	// Model is an interface to be customized, add more methods here,
	// and implement the added methods in customCompanyModel.
	Model interface {
		companyModel
	}

	customCompanyModel struct {
		*defaultCompanyModel
	}
)

// NewCompanyModel returns a model for the database table.
func NewCompanyModel(conn sqlx.SqlConn) Model {
	return &customCompanyModel{
		defaultCompanyModel: newCompanyModel(conn),
	}
}
