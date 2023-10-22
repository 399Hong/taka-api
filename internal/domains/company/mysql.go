package company

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"taka-api/internal/model/mysql/company"
)

type MySqlRepo struct {
	company.Model
}

func NewMySqlRepo(conn sqlx.SqlConn) *MySqlRepo {
	return &MySqlRepo{company.NewCompanyModel(conn)}
}

func (m MySqlRepo) Register(ctx context.Context, register *Register) (*Register, error) {
	registerUid := register.HeadUser.HeadUserID
	if registerUid == 0 {
		return nil, errors.New("no associated user id specified, unable to register company")
	}
	//TODO check whether the user have registered company before
	// if yes return error
	newCompany := company.Company{
		Name:               register.CompanyName,
		Description:        sql.NullString{String: register.CompanyDesc, Valid: true},
		Industry:           register.Industry,
		SizeClassification: register.CompanyScale,
		HeadUserId:         registerUid,
	}
	res, err := m.Insert(ctx, &newCompany)
	if err != nil {
		return nil, errorx.Wrap(err, "unable to create company due to DB error")
	}
	// todo update user name
	_ = res

	return nil, nil
}

func (m MySqlRepo) GetByCompanyId(ctx context.Context, company *Company) (*Company, error) {
	companyId := company.CompanyId
	if companyId == 0 {
		return nil, errors.New("no company id specified")
	}

	res, err := m.FindOne(ctx, companyId)
	if err != nil {
		return nil, errorx.Wrap(err, "unable to find company due to DB error")
	}

	return &Company{
		CompanyId:    res.Id,
		Industry:     res.Industry,
		CompanyScale: res.SizeClassification,
		CompanyValue: 0, /// not implemented yet
		CompanyName:  res.Name,
		CompanyDesc:  res.Description.String,
	}, nil
}

func (m MySqlRepo) GetByHeadUserId(ctx context.Context, headUser *HeadUser) (*Company, error) {
	headUserId := headUser.HeadUserID
	if headUserId == 0 {
		return nil, errors.New("no head user id specified")
	}

	res, err := m.FindOneByHeadUserId(ctx, headUserId)
	if err != nil {
		return nil, errorx.Wrap(err, "unable to find company due to DB error")
	}

	return &Company{
		CompanyId:    res.Id,
		Industry:     res.Industry,
		CompanyScale: res.SizeClassification,
		CompanyValue: 0, /// not implemented yet
		CompanyName:  res.Name,
		CompanyDesc:  res.Description.String,
	}, nil
}
