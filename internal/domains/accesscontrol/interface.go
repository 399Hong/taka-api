package accesscontrol

import "context"

type (
	SsoAuthenticator interface {
		//todo make need to refactor to return info instead of using pointer of original accesor
		//to pass value, not sure which is a better way
		Authenticate(ctx context.Context, method *Accessor) (string, error)
	}

	UserRepo interface {
		Add(ctx context.Context, accessor *Accessor) (*Accessor, error)
		GetById(ctx context.Context, id int64) (*User, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
	}
)
