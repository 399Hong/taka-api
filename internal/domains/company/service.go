package company

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Config func(s *Service) error

type Service struct {
	Repo
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
		s.Repo = NewMySqlRepo(conn)
		return nil
	}
}

//
//func (s *Service) Register(ctx context.Context, register Register) (*Register, error) {
//	user, err := s.Repo.Register(ctx, &register)
//	if err != nil && err != sqlc.ErrNotFound {
//		return nil, errorx.Wrap(err, "unable to get user by email due to DB error")
//	}
//	if user != nil {
//		return nil, errors.New("user is already registered")
//	}
//
//	newAccessor, err = s.Add(ctx, newAccessor)
//	if err != nil {
//		return nil, errorx.Wrap(err, "unable to create user due to DB error")
//	}
//
//	return newAccessor, nil
//}
//
//func (s *Service) LogIn(ctx context.Context, method AccessMethod) (*Accessor, error) {
//	newAccessor, err := s.authenticateAndCreateAccessor(ctx, method)
//	if err != nil {
//		return nil, errorx.Wrap(err, "unable to log in user due to authentication step failure")
//	}
//
//	isPasswordAuth := method.AccessOption == Password
//
//	user, err := s.GetByEmail(ctx, newAccessor.Email)
//	if user.Email != newAccessor.Email {
//		// should not happen
//		return nil, errors.New("email doesn't match")
//	}
//	if isPasswordAuth && user.Password != newAccessor.Password {
//		return nil, errors.New("wrong email or password")
//	}
//
//	newAccessor.SetId(user.Id)
//	return newAccessor, nil
//}
//
//func (s *Service) authenticateAndCreateAccessor(ctx context.Context, method AccessMethod) (*Accessor, error) {
//	newAccessor, err := NewAccessor(method)
//	if err != nil {
//		return nil, err
//	}
//	isPasswordAuth := method.AccessOption == Password
//	if !isPasswordAuth {
//		email, err := s.SsoAuthenticator.Authenticate(ctx, newAccessor)
//		if err != nil {
//			// overwrite email to empty just be safe
//			newAccessor.Email = ""
//			return nil, err
//		}
//		newAccessor.Email = email
//	}
//
//	//TODO need to add proper validation and even send user confirmation email
//	// if its not using SSO to retrieve email
//	if len(newAccessor.Email) == 0 {
//		return nil, errors.New("invalid email address")
//	}
//	//TODO more validation logic
//	if isPasswordAuth && len(newAccessor.Password) == 0 {
//		return nil, errors.New("invalid password")
//	}
//
//	//TODO add salt or other mechanism
//	//TODO what SSO registration, how to set password
//	if isPasswordAuth {
//		newAccessor.Password = hash.Md5Hex([]byte(newAccessor.Password))
//	}
//
//	return newAccessor, nil
//}
