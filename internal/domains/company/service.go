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

func (s *Service) Register(ctx context.Context, register Register) (*Register, error) {
	_, err := s.CompanyRepo.Register(ctx, &register)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, errorx.Wrap(err, "unable to register company for the user")
	}
	//TODO no need to response with cmp info
	return nil, nil
}
