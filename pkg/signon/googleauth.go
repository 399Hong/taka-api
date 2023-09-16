package signon

import (
	"context"
	"errors"
	"google.golang.org/api/idtoken"
)

type GoogleAuth struct{}

func (ga *GoogleAuth) Authenticate(ctx context.Context, token string) (*Claim, error) {
	payload, err := idtoken.Validate(ctx, token, "") //TODO need to validate audience. not sure how audience are set
	if err != nil {
		return nil, err
	}

	claims := payload.Claims
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return nil, errors.New("email not found, unable to identify the user")
	}
	verified, ok := claims["email_verified"].(bool)
	if !ok && !verified {
		return nil, errors.New("email not verified, unable to sign up")
	}

	return &Claim{
		Id:    "",
		Email: email,
	}, nil
}
