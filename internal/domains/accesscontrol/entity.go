package accesscontrol

import (
	"errors"
)

type AccessOption int

// TODO unify with pkg enum
const (
	Password AccessOption = iota + 1
	Google
)

type (
	User struct {
		Id       int64
		Password string
		Email    string
	}

	Tokens struct {
		AccessToken  string
		RefreshToken string
	}

	AccessMethod struct {
		AccessOption
		Token    string
		Email    string
		Password string
	}

	Accessor struct {
		user *User
		AccessMethod
		Tokens
	}
)

func NewAccessor(method AccessMethod) (*Accessor, error) {
	switch method.AccessOption {
	case Password:
	case Google:
	default:
		return nil, errors.New("access option not found")
	}
	user := User{}

	return &Accessor{
		user:         &user,
		AccessMethod: method,
	}, nil
}

func (a *Accessor) GetId() (int64, error) {
	if a.user == nil {
		return 0, errors.New("no user found")
	}
	if a.user.Id == 0 {
		return 0, errors.New("user ID is empty")
	}
	return a.user.Id, nil
}

func (a *Accessor) SetId(id int64) {
	if a.user == nil {
		a.user = &User{Id: id}
	}
	a.user.Id = id
}

func (a *Accessor) SetAccessToken(token string) error {
	if a.user == nil {
		return errors.New("no user found, unable to set access token")
	}

	_, err := a.GetId()
	if err != nil {
		return err
	}

	a.AccessToken = token
	return nil
}

func (a *Accessor) GetAccessToken() (string, error) {
	if a.user == nil {
		return "", errors.New("no user found, unable to get access token")
	}

	if len(a.AccessToken) == 0 {
		return "", errors.New("no access token found")
	}

	return a.AccessToken, nil
}

func (a *Accessor) SetRefreshToken(token string) error {
	if a.user == nil {
		return errors.New("no user found, unable to set refresh token")
	}

	_, err := a.GetId()
	if err != nil {
		return err
	}

	a.RefreshToken = token
	return nil
}

func (a *Accessor) GetRefreshToken() (string, error) {
	if a.user == nil {
		return "", errors.New("no user found, unable to get refresh token")
	}

	if len(a.RefreshToken) == 0 {
		return "", errors.New("no refresh token found")
	}

	return a.RefreshToken, nil
}
