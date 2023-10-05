package singlesignon

import (
	"context"
	"errors"
)

type Authenticator interface {
	// Authenticate will only claim with populated claim, or error will be returned
	Authenticate(ctx context.Context, token string) (*Claim, error)
}

func NewFactory(t AuthType) (Authenticator, error) {
	switch t {
	case Google:
		return &GoogleAuth{}, nil
	case Password:
		return nil, ErrCannotUsePasswordAuth
	}
	return nil, errors.New("no sign-on authentication type found")
}
