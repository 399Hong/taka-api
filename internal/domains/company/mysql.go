package company

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"taka-api/internal/model/mysql/company"
	"taka-api/internal/model/mysql/user"
)

type MySqlRepo struct {
	company.CompanyModel
	user.UserModel
}

func NewMySqlRepo(conn sqlx.SqlConn) *MySqlRepo {
	return &MySqlRepo{
		company.NewCompanyModel(conn),
		user.NewUserModel(conn),
	}
}

func (m *MySqlRepo) Register(ctx context.Context, register *Register) (*Register, error) {
	registerUid := register.HeadUser.HeadUserID
	if registerUid == 0 {
		return nil, errors.New("no associated user id specified, unable to register company")
	}

	newCompany := company.Company{
		Name:               register.CompanyName,
		Description:        sql.NullString{String: register.CompanyDesc, Valid: true},
		Industry:           register.Industry,
		SizeClassification: register.CompanyScale,
		HeadUserId:         registerUid,
	}
	_, err := m.CompanyModel.Insert(ctx, &newCompany)
	if err != nil {
		return nil, errorx.Wrap(err, "unable to create company due to DB error")
	}

	return nil, nil
}

func (m *MySqlRepo) GetByCompanyId(ctx context.Context, companyId int64) (*Company, error) {
	if companyId == 0 {
		return nil, errors.New("no company id specified")
	}

	res, err := m.CompanyModel.FindOne(ctx, companyId)
	if err != nil {
		return nil, errorx.Wrap(err, "unable to find company due to DB error")
	}

	return &Company{
		CompanyId:    res.Id,
		Industry:     res.Industry,
		CompanyScale: res.SizeClassification,
		CompanyValue: 0, //TODO not implemented yet
		CompanyName:  res.Name,
		CompanyDesc:  res.Description.String,
	}, nil
}

func (m *MySqlRepo) GetByHeadUserId(ctx context.Context, headUserId int64) (*Company, error) {
	if headUserId == 0 {
		return nil, errors.New("no head user id specified")
	}

	res, err := m.CompanyModel.FindOneByHeadUserId(ctx, headUserId)
	if err != nil && err == sqlc.ErrNotFound {
		return nil, err
	}
	if err != nil {
		return nil, errorx.Wrap(err, "unable to find company due to DB error")
	}

	return &Company{
		CompanyId:    res.Id,
		Industry:     res.Industry,
		CompanyScale: res.SizeClassification,
		CompanyValue: 0, //TODO not implemented yet
		CompanyName:  res.Name,
		CompanyDesc:  res.Description.String,
	}, nil
}

func (m *MySqlRepo) UpdateCompanyHeadUserInfo(ctx context.Context, headUser *HeadUser) error {
	if headUser == nil {
		return errors.New("no user information specified")
	}
	if headUser.HeadUserID == 0 || headUser.LastName == "" || headUser.FistName == "" || headUser.PhoneNumber == "" {
		return fmt.Errorf("head user infomation is not fully specified %v", headUser)
	}

	_, err := m.CompanyModel.FindOneByHeadUserId(ctx, headUser.HeadUserID)
	if err != nil && err == sqlc.ErrNotFound {
		return fmt.Errorf("the user %v does not belong to any company", headUser.HeadUserID)
	}
	if err != nil {
		return errorx.Wrap(err, "unable to find the company that the user belong to due to DB error")
	}

	err = m.UserModel.UpdateUserNameAndMobile(ctx, &user.User{
		Id:        headUser.HeadUserID,
		FirstName: headUser.FistName,
		LastName:  headUser.LastName,
		Mobile:    headUser.PhoneNumber,
	})
	if err != nil {
		return errorx.Wrap(err, "unable to update head user info")
	}

	return nil
}
