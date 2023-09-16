package signon

import (
	"context"
	"errors"
)

type Authenticator interface {
	Authenticate(ctx context.Context, token string) (*Claim, error)
}

func NewFactory(t AuthType) (Authenticator, error) {
	if t == Google {
		return &GoogleAuth{}, nil
	}
	return nil, errors.New("no sign-on authentication type found")
}
