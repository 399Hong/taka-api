package accesscontroinfra

import (
	"context"
	"taka-api/internal/domains/accesscontrol"
	"taka-api/pkg/singlesignon"
)

type ssoAuthenticator struct{}

func NewSsoAuthenticator() *ssoAuthenticator {
	return &ssoAuthenticator{}
}
func (a *ssoAuthenticator) Authenticate(ctx context.Context, accessor *accesscontrol.Accessor) (string, error) {
	//TODO remove hardcoded token
	accessor.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjgzOGMwNmM2MjA0NmMyZDk0OGFmZmUxMzdkZDUzMTAxMjlmNGQ1ZDEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI0MDc0MDg3MTgxOTIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0MDc0MDg3MTgxOTIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDQ2NTc1NDM5Mjc2NTc3OTg2MTEiLCJoZCI6ImF1Y2tsYW5kdW5pLmFjLm56IiwiZW1haWwiOiJoemhhNzg0QGF1Y2tsYW5kdW5pLmFjLm56IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJWY2J1TXVyWEhNa3ZqT05Sa2FLSjZnIiwiaWF0IjoxNjk0MjQ0NDUyLCJleHAiOjE2OTQyNDgwNTJ9.B23zhpubnTBBxWfPrXjc6KedbMqTveu7jTh5kMRAtYPjZ5XrXYADEOJN5ykvlK0ANipbXHwbOgxcpGzJKwLSWqvnJmOE-yHWxeyJvbsQDn-PTldtDQBSjf8GOcNGmtUTiZj9_ZDBjsZxULmaaaWakEXYiH4QCnY2zPELCwWEZeDRe4ZI5cjl21xvsdAJUfipc_j_7MP2y71wlS5TjeMnEDSX2t1L8SgqohMBM0jCYGE-3Qhaf7WhJWSXY59uBz2iz5k4jXh24vME6DSSxIUa-H1ZdZ8VsEdNzJBRbXKaZa96iyA0j83f3UaGHYvvflYNmpXA0rn55tkJ7dIR4uTgpA"
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
