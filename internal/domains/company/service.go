package company

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Config func(s *Service) error

type Service struct {
	CompanyRepo
}

func NewService(cfgs ...Config) (*Service, error) {
	s := &Service{}

	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithMemoryCompanyRepository() Config {
	return func(s *Service) error {
		return errors.New("not implemented")
	}
}

func WithMySqlCompanyRepo(conn sqlx.SqlConn) Config {
	return func(s *Service) error {
		s.CompanyRepo = NewMySqlRepo(conn)
		return nil
	}
}

func (s *Service) Register(ctx context.Context, register *Register) (*Register, error) {
	if register == nil {
		return nil, errors.New("no register is set")
	}

	company, err := s.GetByHeadUserId(ctx, register.HeadUserID)
	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		return nil, errorx.Wrap(err, "unable to register a company for the user")
	}
	if company != nil {
		return nil, errors.New("user has already registered a company")
	}

	register, err = s.Register(ctx, register)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, errorx.Wrap(err, "unable to register company for the user")
	}

	if err != nil {
		return nil, errorx.Wrap(err, "unable to register company for the user")
	}

	err = s.UpdateCompanyHeadUserInfo(ctx, &register.HeadUser)
	if err != nil {
		// TODO wrap error
		return nil, err
	}
	return register, nil
}
