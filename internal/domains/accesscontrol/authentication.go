package accesscontrol

import (
	"context"
	"taka-api/pkg/singlesignon"
)

type ssoAuthenticator struct{}

func NewSsoAuthenticator() *ssoAuthenticator {
	return &ssoAuthenticator{}
}
func (a *ssoAuthenticator) Authenticate(ctx context.Context, accessor *Accessor) (string, error) {
	auth, err := singlesignon.NewFactory(singlesignon.AuthType(accessor.AccessOption))
	if err != nil {
		return "", err
	}

	claim, err := auth.Authenticate(ctx, accessor.Token)
	if err != nil {
		return "", err
	}

	return claim.Email, nil
}
